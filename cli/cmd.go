package cli

import (
	"os"

	"github.com/fatih/color"
	"github.com/spf13/cobra"

	"github.com/fornellas/prometheus-mdns-http-sd/cli/discover"
	"github.com/fornellas/prometheus-mdns-http-sd/cli/lib"
	"github.com/fornellas/prometheus-mdns-http-sd/cli/server"
	"github.com/fornellas/prometheus-mdns-http-sd/cli/version"
	"github.com/fornellas/prometheus-mdns-http-sd/log"
)

var ExitFunc func(int) = func(code int) { os.Exit(code) }

var logLevelStr string
var defaultLogLevelStr = "info"
var forceColor bool
var defaultForceColor = false

var Cmd = &cobra.Command{
	Use:   "prometheus-mdns-http-sd",
	Short: "Prometheus HTTP mDNS Service Discovery.",
	Args:  cobra.NoArgs,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		if forceColor {
			color.NoColor = false
		}
		cmd.SetContext(log.SetLoggerValue(
			cmd.Context(), cmd.OutOrStderr(), logLevelStr, ExitFunc,
		))
	},
	Run: func(cmd *cobra.Command, args []string) {
		logger := log.GetLogger(cmd.Context())
		if err := cmd.Help(); err != nil {
			logger.Fatal(err)
		}
	},
}

var resetFuncs []func()

func Reset() {
	logLevelStr = defaultLogLevelStr
	forceColor = defaultForceColor
	for _, resetFunc := range resetFuncs {
		resetFunc()
	}
	lib.Reset()
}

func init() {
	Cmd.PersistentFlags().StringVarP(
		&logLevelStr, "log-level", "l", defaultLogLevelStr,
		"Logging level",
	)
	Cmd.PersistentFlags().BoolVarP(
		&forceColor, "force-color", "", defaultForceColor,
		"Force colored output",
	)

	Cmd.AddCommand(discover.Cmd)
	resetFuncs = append(resetFuncs, discover.Reset)

	Cmd.AddCommand(server.Cmd)
	resetFuncs = append(resetFuncs, server.Reset)

	Cmd.AddCommand(version.Cmd)
	resetFuncs = append(resetFuncs, version.Reset)
}
