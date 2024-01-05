package main

import (
	"log"

	"github.com/fornellas/prometheus-mdns-http-sd/cli"
)

func main() {
	log.SetFlags(0)
	if err := cli.Cmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
