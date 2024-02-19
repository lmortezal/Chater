package cmd
import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var port int
var domain string
var server bool = false

var rootCmd = &cobra.Command{
	Use:   "Chater",
	Short: "This is a sample Chat application",
}

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Run in server mode",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Running in server mode on port %d and domain %s\n", port, domain)
		server = true
	},
}

var clientCmd = &cobra.Command{
	Use:   "client",
	Short: "Run in client mode",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Running in client mode on port %d and domain %s\n", port, domain)
		server = false
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)
	rootCmd.AddCommand(clientCmd)
	
	rootCmd.PersistentFlags().StringVarP(&domain, "domain", "d", "localhost", "Domain name")
	rootCmd.PersistentFlags().IntVarP(&port, "port", "p", 8080, "Port number")
}

func Execute() (string, int , bool){
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return domain, port, server
}