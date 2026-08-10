package main

import (
	"log"
	"os"

	"github.com/lmortezal/Chater/internal/server"
	"github.com/spf13/cobra"
)

func main() {
	var port int
	var domain string

	var rootCmd = &cobra.Command{
		Use:          "chater-server",
		Short:        "Run Chater in server mode",
		SilenceUsage: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Printf("Running in server mode on port %d and domain %s\n", port, domain)
			return server.Startlistening(domain, port)
		},
	}
	rootCmd.Flags().StringVarP(&domain, "domain", "d", "localhost", "Domain name")
	rootCmd.Flags().IntVarP(&port, "port", "p", 8080, "Port number")
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
