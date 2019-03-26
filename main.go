package main

import (
	log "github.com/sirupsen/logrus"
	"github.com/tarkalabs/artisanal-containers/cmd"
)

func init() {
	log.SetLevel(log.DebugLevel)
}

func main() {
	cmd.Execute()
}
