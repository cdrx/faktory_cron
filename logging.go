package main

import (
	"os"
	"time"

	apex "github.com/apex/log"
	"github.com/apex/log/handlers/text"
)

var log *apex.Entry

func configureLogging() {
	apex.SetHandler(text.New(os.Stderr))
	log = apex.WithFields(apex.Fields{
		"time": time.Now().Format(time.RFC3339),
	})
}
