package server

import (
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/spf13/cobra"

	"github.com/fornellas/prometheus-mdns-http-sd/cli/lib"
	"github.com/fornellas/prometheus-mdns-http-sd/log"
	"github.com/fornellas/prometheus-mdns-http-sd/server"
)

var addr string

var Cmd = &cobra.Command{
	Use:   "server",
	Short: "Start a server suitable for Prometheus http_sd_config.",
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		ctx := cmd.Context()

		logger := log.GetLogger(ctx)

		srv := server.NewServer(
			addr,
			lib.InterfaceStr,
			lib.Service,
			lib.Domain,
			lib.Timeout,
			lib.DisableIPv4,
			lib.DisableIPv6,
		)

		go func() {
			sig := make(chan os.Signal, 1)
			signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
			<-sig

			logger.Info("Shutting down...")
			if err := srv.Shutdown(ctx); err != nil {
				logger.Errorf("Shutdown request failed: %v", err)
			}
		}()

		logger.Infof("Starting server on %s", addr)
		if err := srv.ListenAndServe(); err != http.ErrServerClosed {
			logger.Fatalf("Server error: %v", err)
		}
		logger.Info("Exiting")
	},
}

func Reset() {
}

func init() {
	lib.AddCommonFlags(Cmd)

	Cmd.Flags().StringVarP(
		&addr, "address", "", ":2431",
		"TCP address for the server to listen on.",
	)
}
