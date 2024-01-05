package discover

import (
	"fmt"
	"net"
	"time"

	"github.com/hashicorp/mdns"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var defaultService = "_prometheus-http._tcp"
var service string

var defaultDomain = "local"
var domain string

var defaultTimeout = time.Second
var timeout time.Duration

var defaultIntterfaceStr = "all"
var interfaceStr string

var defaultWantUnicastResponse = false
var wantUnicastResponse bool

var defaultDisableIPv4 = false
var disableIPv4 bool

var defaultDisableIPv6 = false
var disableIPv6 bool

var Cmd = &cobra.Command{
	Use:   "discover",
	Short: "Runs mDNS discovery and print results.",
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		entriesCh := make(chan *mdns.ServiceEntry, 4)
		go func() {
			for entry := range entriesCh {
				fmt.Printf(
					"%s\n  Host: %s\n  AddrV4: %v\n  AddrV6: %v\n  Port: %v\n  Info: %v\n  InfoFields:\n",
					entry.Name,
					entry.Host,
					entry.AddrV4,
					entry.AddrV6,
					entry.Port,
					entry.Info,
				)
				for _, infoField := range entry.InfoFields {
					fmt.Printf("    %v\n", infoField)
				}
			}
		}()
		defer close(entriesCh)

		var netInterface *net.Interface
		if interfaceStr != defaultIntterfaceStr {
			var err error
			netInterface, err = net.InterfaceByName(interfaceStr)
			if err != nil {
				logrus.Fatalf("Invalid interface '%s': %v", interfaceStr, err)
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
			logrus.Fatalf("Error querying mDNS: %v", err)
		}

	},
}

func Reset() {
	service = defaultService
	domain = defaultDomain
	timeout = defaultTimeout
}

func init() {
	Cmd.PersistentFlags().StringVarP(
		&service, "service", "s", defaultService,
		"Service",
	)

	Cmd.PersistentFlags().StringVarP(
		&domain, "domain", "d", defaultDomain,
		"Domain",
	)

	Cmd.PersistentFlags().DurationVarP(
		&timeout, "timeout", "t", defaultTimeout,
		"Timeout",
	)

	Cmd.PersistentFlags().StringVarP(
		&interfaceStr, "interface", "i", defaultIntterfaceStr,
		"Multicast interface to use",
	)

	Cmd.PersistentFlags().BoolVarP(
		&wantUnicastResponse, "want-unicast-response", "w", defaultWantUnicastResponse,
		"Unicast response desired, as per 5.4 in RFC",
	)

	Cmd.PersistentFlags().BoolVarP(
		&disableIPv4, "disable-ipv4", "", defaultDisableIPv4,
		"Whether to disable usage of IPv4 for MDNS operations. Does not affect discovered addresses.",
	)

	Cmd.PersistentFlags().BoolVarP(
		&disableIPv6, "disable-ipv6", "", defaultDisableIPv6,
		"Whether to disable usage of IPv6 for MDNS operations. Does not affect discovered addresses.",
	)
}
