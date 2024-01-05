package discover

import (
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "discover",
	Short: "Runs mDNS discovery and print results.",
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func Reset() {

}

func init() {

}
