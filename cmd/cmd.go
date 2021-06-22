package cmd

import (
	"os"

	"github.com/rige1/sberapi-mock/config"
	"github.com/rige1/sberapi-mock/server"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
		port int32 
		cert string
		key string
		clientCert string
		ignoreValidation bool

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
	startCmd.PersistentFlags().StringVar(&cert, "cert", "", "server certificate")
	startCmd.PersistentFlags().StringVar(&key, "key", "", "server certificate key")
	startCmd.PersistentFlags().StringVar(&clientCert, "client-cert", "", "client certificate (mTLS) for verification")
	startCmd.PersistentFlags().BoolVar(&ignoreValidation, "ignore-validation", false, "ignore validation of incoming request, default false")
	rootCmd.AddCommand(startCmd)
}


func start(cmd *cobra.Command, args[]string) {

	config := config.Config {
		Port: int(port),
		Cert: cert,
		Key: key,
		ClientCert: clientCert,
		Validate: !ignoreValidation,
	}

	mockServer, err := server.NewMockServerDefault(config)

	if err != nil {
		log.Error(err)
		return
	}
	
	err = mockServer.Run()

	if err != nil {
		log.Error(err)
		return
	}
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Error(err)
		os.Exit(1)
	}
}