package cmd

import (
	"log"
	"os"
	"os/signal"

	"github.com/spf13/cobra"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "start the service",
	Long:  "start the service with the give service name",
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("read the cfgfile and start the config name")
		waitClosed := make(chan struct{})
		go func() {
			sigint := make(chan os.Signal, 1)
			signal.Notify(sigint, os.Interrupt)
			<-sigint
			//do the close procedure here.
			close(waitClosed)
		}()
		//start the service, typically block here
		<-waitClosed
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}
