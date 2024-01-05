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
		entries, err := mdns.Query(
			lib.InterfaceStr,
			lib.Service,
			lib.Domain,
			lib.Timeout,
			lib.WantUnicastResponse,
			lib.DisableIPv4,
			lib.DisableIPv6,
		)
		if err != nil {
			logrus.Fatalf("Error querying mDNS: %v", err)
		}

		for _, entry := range entries {
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
	},
}

func Reset() {
}

func init() {
	lib.AddCommonFlags(Cmd)
}
