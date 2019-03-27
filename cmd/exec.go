package cmd

import (
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/tarkalabs/artisanal-containers/container"
)

var execCommand = &cobra.Command{
	Use: "exec",
	Run: func(cmd *cobra.Command, args []string) {
		var otherArgs []string
		if len(args) > 1 {
			otherArgs = args[1:]
		}
		ac := &container.ArtisanalContainer{Command: args[0], Args: otherArgs, Uid: os.Getuid(), Gid: os.Getgid()}
		err := ac.Start(isFork)
		if err != nil {
			log.Fatal(err)
		}
	},
}
var isFork bool

func init() {
	execCommand.Flags().BoolVar(&isFork, "fork", false, "--fork=true (used internally)")
}
