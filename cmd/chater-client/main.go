package main

import (
	"fmt"
	"os"

	tui "github.com/lmortezal/Chater/internal/client/tui"
	"github.com/spf13/cobra"
)

var port int
var domain string

var rootCmd = &cobra.Command{
	Use:          "chater-client",
	Short:        "Run Chater in client mode",
	SilenceUsage: true,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Running in client mode on port %d and domain %s\n", port, domain)
		tui.Tui_main(domain, port)
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
