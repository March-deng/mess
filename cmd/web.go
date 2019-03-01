package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

var webCmd = &cobra.Command{
	Use:   "web",
	Short: "start the webService",
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("start the web service")
	},
}

func init() {
	serveCmd.AddCommand(webCmd)
}
