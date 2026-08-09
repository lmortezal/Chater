package main

import (
	"fmt"
	"os"

	"github.com/lmortezal/Chater/internal/server"
	"github.com/spf13/cobra"
)

var port int
var domain string

var rootCmd = &cobra.Command{
	Use:          "chater-server",
	Short:        "Run Chater in server mode",
	SilenceUsage: true,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Running in server mode on port %d and domain %s\n", port, domain)
		server.Startlistening(domain, port)
	},
}

func main() {
	rootCmd.Flags().StringVarP(&domain, "domain", "d", "localhost", "Domain name")
	rootCmd.Flags().IntVarP(&port, "port", "p", 8080, "Port number")
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
