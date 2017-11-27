package main

import (
	"os"

	apex "github.com/apex/log"
	"github.com/apex/log/handlers/cli"
)

var log *apex.Entry

func configureLogging() {
	apex.SetLevel(apex.DebugLevel)
	handler := cli.New(os.Stdout)
	handler.Padding = 0
	apex.SetHandler(handler)
	log = apex.WithFields(apex.Fields{})
}
