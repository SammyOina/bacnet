package bacnet

import (
	"context"
	"errors"

	"github.com/absmach/bacnet/pkg/bacnet"
	"github.com/absmach/bacnet/pkg/encoding"
	"github.com/absmach/bacnet/pkg/transport"
	"golang.org/x/sync/errgroup"
)

var _ Client = (*client)(nil)

var errNoResponse = errors.New("no response received")

type Client interface {
	ReadProperty(ctx context.Context, address string, request bacnet.ReadPropertyRequest) ([]bacnet.BACnetValue, error)
	WriteProperty(ctx context.Context, address string, request bacnet.WritePropertyRequest) error
}

type client struct {
	transport transport.Transport
}

// ReadProperty implements Client.
func (c *client) ReadProperty(ctx context.Context, address string, request bacnet.ReadPropertyRequest) ([]bacnet.BACnetValue, error) {
	destination, err := bacnet.NewBACnetAddress(0, nil, address)
	if err != nil {
		return []bacnet.BACnetValue{}, err
	}
	npdu := bacnet.NewNPDU(destination, nil, nil, nil)
	npdu.Control.SetDataExpectingReply(true)
	npdu.Control.SetNetworkPriority(bacnet.NormalMessage)

	npduBytes, err := npdu.Encode()
	if err != nil {
		return []bacnet.BACnetValue{}, err
	}

	apdu := bacnet.APDU{
		PduType:                   bacnet.PDUTypeConfirmedServiceRequest,
		ServiceChoice:             byte(bacnet.ReadProperty),
		SegmentedResponseAccepted: false,
		MaxSegmentsAccepted:       bacnet.BacnetMaxSegments(encoding.NoSegmentation),
		InvokeID:                  0,
	}

	mes := append(npduBytes, apdu.Encode()...)
	mes = append(mes, request.Encode()...)

	res := make(chan []byte, 1)
	var eg errgroup.Group

	eg.Go(func() error {
		defer close(res)
		return c.transport.Send(ctx, address, mes, int(bacnet.BVLCOriginalBroadcastNPDU), res)
	})

	// Wait for all goroutines to finish using the error group
	if err := eg.Wait(); err != nil {
		return []bacnet.BACnetValue{}, err
	}

	select {
	case <-ctx.Done():
		return []bacnet.BACnetValue{}, ctx.Err()
	case response, ok := <-res:
		if !ok {
			return []bacnet.BACnetValue{}, errNoResponse
		} else {
			blvc := bacnet.BVLC{BVLLTypeBACnetIP: 0x81}
			headerLength, _, _, err := blvc.Decode(response, 0)
			if err != nil {
				return []bacnet.BACnetValue{}, err
			}
			npduRes := bacnet.NPDU{Version: 1}
			npduLen, err := npduRes.Decode(response, headerLength)
			if err != nil {
				return []bacnet.BACnetValue{}, err
			}
			apduRes := bacnet.APDU{}
			apduLen := apduRes.Decode(response, headerLength+npduLen)
			readPropACK := bacnet.ReadPropertyACK{}
			if _, err = readPropACK.Decode(response, headerLength+npduLen+apduLen-2, len(response)); err != nil {
				return []bacnet.BACnetValue{}, err
			}
			return readPropACK.PropertyValue, nil
		}
	default:
		return []bacnet.BACnetValue{}, errNoResponse
	}
}

// WriteProperty implements Client.
func (c *client) WriteProperty(ctx context.Context, address string, request bacnet.WritePropertyRequest) error {
	destination, err := bacnet.NewBACnetAddress(0, nil, "127.0.0.6:47809")
	if err != nil {
		return err
	}
	npdu := bacnet.NewNPDU(destination, nil, nil, nil)
	npdu.Control.SetDataExpectingReply(true)
	npdu.Control.SetNetworkPriority(bacnet.NormalMessage)

	npduBytes, err := npdu.Encode()
	if err != nil {
		return err
	}

	apdu := bacnet.APDU{
		PduType:                   bacnet.PDUTypeConfirmedServiceRequest,
		ServiceChoice:             byte(bacnet.WriteProperty),
		SegmentedResponseAccepted: false,
		MaxSegmentsAccepted:       bacnet.BacnetMaxSegments(encoding.NoSegmentation),
		InvokeID:                  0,
	}

	mes := append(npduBytes, apdu.Encode()...)
	mes = append(mes, request.Encode()...)

	res := make(chan []byte, 1)
	var eg errgroup.Group

	eg.Go(func() error {
		defer close(res)
		return c.transport.Send(ctx, address, mes, int(bacnet.BVLCOriginalBroadcastNPDU), res)
	})

	// Wait for all goroutines to finish using the error group
	if err := eg.Wait(); err != nil {
		return err
	}
	select {
	case <-ctx.Done():
		return ctx.Err()
	case response, ok := <-res:
		if !ok {
			return errNoResponse
		} else {
			blvc := bacnet.BVLC{BVLLTypeBACnetIP: 0x81}
			headerLength, _, _, err := blvc.Decode(response, 0)
			if err != nil {
				return err
			}
			npduRes := bacnet.NPDU{Version: 1}
			npduLen, err := npduRes.Decode(response, headerLength)
			if err != nil {
				return err
			}
			apduRes := bacnet.APDU{}
			_ = apduRes.Decode(response, headerLength+npduLen)
			return nil
		}
	default:
		return errNoResponse
	}
}