package lib

import (
	"time"

	"github.com/spf13/cobra"

	"github.com/fornellas/prometheus-mdns-http-sd/mdns"
)

var DefaultService = "_prometheus-http._tcp"
var Service string

var DefaultDomain = "local"
var Domain string

var DefaultTimeout = time.Second
var Timeout time.Duration

var DefaultIntterfaceStr = mdns.AllInterfaces
var InterfaceStr string

var DefaultWantUnicastResponse = false
var WantUnicastResponse bool

var DefaultDisableIPv4 = false
var DisableIPv4 bool

var DefaultDisableIPv6 = false
var DisableIPv6 bool

func AddCommonFlags(cmd *cobra.Command) {
	cmd.PersistentFlags().StringVarP(
		&Service, "service", "s", DefaultService,
		"Service",
	)

	cmd.PersistentFlags().StringVarP(
		&Domain, "domain", "d", DefaultDomain,
		"Domain",
	)

	cmd.PersistentFlags().DurationVarP(
		&Timeout, "timeout", "t", DefaultTimeout,
		"Timeout",
	)

	cmd.PersistentFlags().StringVarP(
		&InterfaceStr, "interface", "i", DefaultIntterfaceStr,
		"Multicast interface to use",
	)

	cmd.PersistentFlags().BoolVarP(
		&WantUnicastResponse, "want-unicast-response", "w", DefaultWantUnicastResponse,
		"Unicast response desired, as per 5.4 in RFC",
	)

	cmd.PersistentFlags().BoolVarP(
		&DisableIPv4, "disable-ipv4", "", DefaultDisableIPv4,
		"Whether to disable usage of IPv4 for MDNS operations. Does not affect discovered addresses.",
	)

	cmd.PersistentFlags().BoolVarP(
		&DisableIPv6, "disable-ipv6", "", DefaultDisableIPv6,
		"Whether to disable usage of IPv6 for MDNS operations. Does not affect discovered addresses.",
	)
}

func Reset() {
	Service = DefaultService
	Domain = DefaultDomain
	Timeout = DefaultTimeout
	InterfaceStr = DefaultIntterfaceStr
	WantUnicastResponse = DefaultWantUnicastResponse
	DisableIPv4 = DefaultDisableIPv4
	DisableIPv6 = DefaultDisableIPv6
}
