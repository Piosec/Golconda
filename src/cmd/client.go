package cmd

import (
	"fmt"
	"github.com/TwinProduction/go-color"
	"github.com/spf13/cobra"
	"golconda/src"
	"golconda/src/client"
	"golconda/src/log"
)

var clientCmd = &cobra.Command{
	Use:   "client",
	Short: "Use client to discover reverse ports.",
	Long: `A Fast client to discover reverse ports made with love.
         It is also possible to get a oneliner in several languages.
         Please use it with love too <3.`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Init(verbose)
		fflags := cmd.Flags()
		if src.ValidateIP(target) {
			if fflags.Changed("ports") {
				if fflags.Changed("cmd") {

					fmt.Println(client.GetClientCommand(target, ports, command))
				} else {
					client.PortRunner(target, ports,excludeports)
				}
			} else if topports > 0 && topports < 65535 {
				if fflags.Changed("cmd") {
					fmt.Println(client.GetClientCommand(target, src.GetTopPorts(topports), command))
				} else {
					client.PortRunner(target, src.GetTopPorts(topports),excludeports)
				}
			} else {
				fmt.Println(color.Ize(color.Red, "[-] ") + "Top ports must be between 1 and 65535")
			}

		} else {
			fmt.Println(color.Ize(color.Red, "[-] ") + "You must provide a correct IPv4 address")
		}

	},
}

func init() {
	clientCmd.Flags().StringVarP(&ports, "ports", "p", "", "Try to initiate a connection to the defined ports. Ex: 21,22,8080 or 22-8080")
	clientCmd.Flags().StringVarP(&target, "target", "t", "127.0.0.1", "Target to check reverse ports. Ex: 127.0.0.1")
	clientCmd.Flags().IntVar(&topports, "top-ports", 100, "Try to initiate a connection to the most known ports.")
	clientCmd.Flags().StringVarP(&command, "cmd", "c", "", "Oneliner in several languages to execute if you cannot upload the binary to the remote target.")
	clientCmd.MarkFlagRequired("target")
	rootCmd.AddCommand(clientCmd)
}
