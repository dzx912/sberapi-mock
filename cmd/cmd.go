package cmd

import (
	"fmt"
	"os"

	"github.com/rige1/sberapi-mock/server"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
		port int32 

		rootCmd = &cobra.Command{
			Use: "sbermock",
			Short: "SberAPI mock server",
		}

		startCmd = &cobra.Command{
			Use: "start",
			Short: "Start mock server",
			Run: start,
		}
)

func init() {
	startCmd.PersistentFlags().Int32Var(&port, "port", 8080, "server port, default 8080")
	rootCmd.AddCommand(startCmd)
}


func start(cmd *cobra.Command, args[]string) {
	mockServer, err := server.NewMockServerDefault()

	if err != nil {
		log.Error(err)
		return
	}
	
	mockServer.Run(fmt.Sprintf(":%d", port))
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}