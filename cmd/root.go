package cmd

import (
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func init() {
	rootCommand.AddCommand(execCommand)
}

var rootCommand = &cobra.Command{
	Use:   "artisanal-containers",
	Short: "Unless you are learning about containers, you should be using docker",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("use help to know more")
	},
}

func Execute() {
	err := rootCommand.Execute()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}
