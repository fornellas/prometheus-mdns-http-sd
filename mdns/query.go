package mdns

import (
	"net"
	"time"

	"github.com/hashicorp/mdns"
)

var AllInterfaces = "all"

func Query(
	interfaceStr string,
	service string,
	domain string,
	timeout time.Duration,
	wantUnicastResponse bool,
	disableIPv4 bool,
	disableIPv6 bool,
) ([]mdns.ServiceEntry, error) {
	var entries []mdns.ServiceEntry

	entriesCh := make(chan *mdns.ServiceEntry, 4)
	go func() {
		for entry := range entriesCh {
			entries = append(entries, *entry)
		}
	}()
	defer close(entriesCh)

	var netInterface *net.Interface
	if interfaceStr != AllInterfaces {
		var err error
		netInterface, err = net.InterfaceByName(interfaceStr)
		if err != nil {
			return nil, err
		}
	}

	err := mdns.Query(
		&mdns.QueryParam{
			Service:             service,
			Domain:              domain,
			Timeout:             timeout,
			Interface:           netInterface,
			Entries:             entriesCh,
			WantUnicastResponse: wantUnicastResponse,
			DisableIPv4:         disableIPv4,
			DisableIPv6:         disableIPv6,
		},
	)
	if err != nil {
		return nil, err
	}

	return entries, nil
}
