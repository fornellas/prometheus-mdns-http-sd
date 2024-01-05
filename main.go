package main

import (
	"log"

	"github.com/fornellas/go_build_template/cli"
)

func main() {
	log.SetFlags(0)
	if err := cli.Cmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
