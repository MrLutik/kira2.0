package main

import (
	"github.com/mrlutik/kira2.0/internal/logging"
	managerCli "github.com/mrlutik/kira2.0/internal/manager/cli"
)

var log = logging.Log

func main() {
	log.Infoln("Starting kira2...")

	managerCli.Start()
}
