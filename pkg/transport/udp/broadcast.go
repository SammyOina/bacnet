package udp

import (
	"net"
	"strconv"

	"github.com/absmach/bacnet/pkg/bacnet"
	"github.com/absmach/bacnet/pkg/encoding"
)

const globalBroadcast = "255.255.255.255"

// GetBroadcastAddress returns the broadcast address given the local address and port.
func GetBroadcastAddress(localEndpoint string, port int) (bacnet.Address, error) {
	broadcast := globalBroadcast
	interfaces, err := net.Interfaces()
	if err != nil {
		return bacnet.Address{}, err
	}

	for _, iface := range interfaces {
		addrs, err := iface.Addrs()
		if err != nil {
			return bacnet.Address{}, err
		}

		for _, addr := range addrs {
			if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() && ipnet.IP.To4() != nil {
				ipAddress := ipnet.IP.String()
				if ipAddress == localEndpoint {
					if iface.Flags&net.FlagBroadcast != 0 {
						broadcast = ipnet.IP.Mask(ipnet.IP.DefaultMask()).String()
					}
				}
			}
		}
	}
	netType := encoding.IPV4
	return bacnet.NewAddress(0xFFFF, nil, broadcast+":"+strconv.Itoa(port), &netType), nil

}
