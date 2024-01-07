package discover

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/fornellas/prometheus-mdns-http-sd/cli/lib"
	"github.com/fornellas/prometheus-mdns-http-sd/mdns"
)

var Cmd = &cobra.Command{
	Use:   "discover",
	Short: "Runs mDNS discovery and print results.",
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		m := mdns.NewMDNS()

		proto := mdns.ProtoAny
		if lib.DisableIPv4 {
			proto = mdns.ProtoInet6
		}
		if lib.DisableIPv6 {
			proto = mdns.ProtoInet
		}

		services, err := m.BrowseServices(
			lib.InterfaceStr,
			proto,
			lib.Service,
			lib.Domain,
			lib.Timeout,
		)
		if err != nil {
			logrus.Fatalf("Error querying mDNS: %v", err)
		}

		for _, service := range services {
			fmt.Printf(
				"%s\n  Type: %v\n  Interface: %v\n  Host: %v\n  Domain: %v\n  IP: %v\n  Port: %v\n  Protocol: %v\n",
				service.Name,
				service.Type,
				service.Interface,
				service.Host,
				service.Domain,
				service.IP,
				service.Port,
				service.Protocol,
			)
		}
	},
}

func Reset() {
}

func init() {
	lib.AddCommonFlags(Cmd)
}
