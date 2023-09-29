package main

import (
	"fmt"
	"log"
	"net"
	"time"

	"github.com/absmach/bacnet/pkg/bacnet"
	"github.com/absmach/bacnet/pkg/encoding"
	"github.com/absmach/bacnet/pkg/transport"
)

func main() {
	serverIp := "127.0.0.6:47809"
	serverLocalAddr := "127.0.0.1:0"
	netType := encoding.IPV4
	destination := bacnet.NewAddress(0, nil, serverIp, &netType)
	npdu := bacnet.NewNPDU(&destination, nil, nil, nil)
	npdu.Control.SetDataExpectingReply(true)
	npdu.Control.SetNetworkPriority(bacnet.NormalMessage)

	npduBytes, err := npdu.Encode()
	if err != nil {
		log.Fatal(err)
	}

	apdu := bacnet.APDU{
		PduType:                   bacnet.PDUTypeConfirmedServiceRequest,
		ServiceChoice:             byte(bacnet.ReadProperty),
		SegmentedResponseAccepted: false,
		MaxSegmentsAccepted:       bacnet.MaxSegments(encoding.NoSegmentation),
		InvokeID:                  0,
	}

	apduBytes, err := apdu.Encode()
	if err != nil {
		log.Fatal(err)
	}

	req := bacnet.ReadPropertyRequest{
		PropertyIdentifier: encoding.PresentValue,
		ObjectIdentifier:   &bacnet.ObjectIdentifier{Type: encoding.AnalogInput, Instance: 10},
	}

	mes := append(npduBytes, apduBytes...)
	mes = append(mes, req.Encode()...)

	blvc, err := bacnet.NewBVLC(transport.IP)
	if err != nil {
		log.Fatal(err)
	}
	blvcBytes := blvc.Encode(bacnet.BVLCOriginalBroadcastNPDU, uint16(len(mes)+4))
	message := append(blvcBytes, mes...)

	remoteAddr, err := net.ResolveUDPAddr("udp", serverIp)
	if err != nil {
		fmt.Println("Error resolving remote address:", err)
		return
	}

	localAddr, err := net.ResolveUDPAddr("udp", serverLocalAddr)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	// Create a UDP connectionBACnetAddress
	conn, err := net.DialUDP("udp", localAddr, remoteAddr)
	if err != nil {
		fmt.Println("Error creating UDP connection:", err)
		return
	}
	defer conn.Close()

	_, err = conn.Write(message)
	if err != nil {
		log.Fatal("Error sending WhoIsRequest:", err)
	}

	// Wait for responses
	buffer := make([]byte, 1500)
	conn.SetReadDeadline(time.Now().Add(5 * time.Second)) // Set a timeout for responses

	for {
		n, _, err := conn.ReadFromUDP(buffer)
		if err != nil {
			if netErr, ok := err.(net.Error); ok && netErr.Timeout() {
				// Timeout reached, no more responses
				log.Println("No more responses received.")
				break
			}
			log.Println("Error reading response:", err)
			break
		}

		response := buffer[:n]
		blvc := bacnet.BVLC{BVLLTypeBACnetIP: blvc.BVLLTypeBACnetIP}
		headerLength, function, msgLength, err := blvc.Decode(response, 0)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("headerLength %v BVLCfunction %v msgLen %v\n", headerLength, function, msgLength)
		fmt.Println("blvc", blvc)
		npdu := bacnet.NPDU{Version: 1}
		npduLen := npdu.Decode(response, headerLength)
		fmt.Println("npdu", npdu)
		apdu := bacnet.APDU{}
		apduLen, err := apdu.Decode(response, headerLength+npduLen)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("apdu", apdu)
		readPropACK := bacnet.ReadPropertyACK{}
		if _, err = readPropACK.Decode(response, headerLength+npduLen+apduLen-2, len(response)); err != nil {
			log.Fatal(err)
		}
		fmt.Println("readprop", readPropACK)
	}
}
