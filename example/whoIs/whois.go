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

	var highLimit, lowLimit uint32 = 4000000, 0
	req := bacnet.WhoIs{
		HighLimit: &highLimit,
		LowLimit:  &lowLimit,
	}
	whoisBytes := req.Encode()

	netType := encoding.IPV4
	broads := bacnet.NewAddress(encoding.MaxUint16-1, nil, serverIp, &netType)

	npdu := bacnet.NewNPDU(&broads, nil, nil, nil)
	npdu.Control.SetNetworkPriority(bacnet.NormalMessage)
	npduBytes, err := npdu.Encode()
	if err != nil {
		log.Fatalf("failed to encode npdu with error %v", err)
	}

	apdu := bacnet.APDU{
		PduType:       bacnet.PDUTypeUnconfirmedServiceRequest,
		ServiceChoice: byte(bacnet.ServiceChoiceWhoIs),
	}

	apduBytes, err := apdu.Encode()
	if err != nil {
		log.Fatal(err)
	}

	mes := append(npduBytes, apduBytes...)
	mes = append(mes, whoisBytes...)

	blvc, err := bacnet.NewBVLC(transport.IP)
	if err != nil {
		log.Fatal(err)
	}

	blvcBytes := blvc.Encode(bacnet.BVLCOriginalBroadcastNPDU, uint16(len(mes)+int(blvc.BVLCHeaderLength)))
	message := append(blvcBytes, mes...)

	remoteAddr, err := net.ResolveUDPAddr("udp", serverIp)
	if err != nil {
		log.Fatal("Error resolving remote address:", err)
	}

	localAddr, err := net.ResolveUDPAddr("udp", serverLocalAddr)
	if err != nil {
		log.Fatal("Error: ", err)
		return
	}

	conn, err := net.DialUDP("udp", localAddr, remoteAddr)
	if err != nil {
		log.Fatal("Error creating UDP connection:", err)
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
		log.Printf("Received response: %X\n", response)
		blvc := bacnet.BVLC{BVLLTypeBACnetIP: blvc.BVLLTypeBACnetIP}
		headerLength, function, msgLength, err := blvc.Decode(response, 0)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(response)
		fmt.Printf("headerLength %v BVLCfunction %v msgLen %v\n", headerLength, function, msgLength)
		fmt.Println("blvc", blvc)
		fmt.Println(response[headerLength:])
		npdu := bacnet.NPDU{Version: 1}
		npduLen := npdu.Decode(response, headerLength)
		fmt.Println("npdu", npdu)
		fmt.Println(response[headerLength+npduLen:])
		apdu := bacnet.APDU{}
		apduLen, err := apdu.Decode(response, headerLength+npduLen)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("apdu", apdu)
		fmt.Println(response[headerLength+npduLen+apduLen:])
		iam := bacnet.IAmRequest{}
		iamLen, err := iam.Decode(response, headerLength+npduLen+apduLen)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("iam", iam)
		fmt.Println(response[headerLength+npduLen+apduLen+iamLen:])
	}
}
