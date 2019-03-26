package cmd

import (
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/tarkalabs/artisanal-containers/container"
)

var rootCommand = &cobra.Command{
	Use:   "artisanal-containers",
	Short: "Unless you are learning about containers, you should be using docker",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(args)
		ac := &container.ArtisanalContainer{Command: args[0], Args: args[1:], Uid: os.Getuid(), Gid: os.Getgid()}
		err := ac.Start()
		if err != nil {
			log.Fatal(err)
		}
	},
}

func Execute() {
	err := rootCommand.Execute()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}
