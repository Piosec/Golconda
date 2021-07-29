package cmd

import (
	"github.com/spf13/cobra"
)

var target string
var ports string
var topports int
var command string
var dump bool
var interfaceName string
var verbose int

var rootCmd = &cobra.Command{
	Use:   "golconda",
	Short: "Golconda is a client/server reverse ports discovery.",
	Long: `A Fast client / server reverse port discovery made with love.
         Please use it with love too <3.`,
}

// Execute get required params depending on the client and server
func Execute() error {
	rootCmd.PersistentFlags().IntVarP(&verbose, "verbose","v", 0, "Increase the verbosity output. (Default: 0)")
	return rootCmd.Execute()
}
