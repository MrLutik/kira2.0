package main

import (
	"github.com/mrlutik/kira2.0/internal/cli"
	"github.com/mrlutik/kira2.0/internal/logging"
)

func main() {
	log := logging.Log
	log.Infoln("Starting kira2launcher...")
	cli.Start()
}
