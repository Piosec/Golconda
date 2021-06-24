package cmd

import (
   "github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
  Use:   "golconda",
  Short: "Golconda is a client/server reverse ports discovery.",
  Long: `A Fast client / server reverse port discovery made with love.
         Please use it with love too <3.`,
}

var target string
var ports string
var topports int
var command string
var dump bool
var interface_name string

func Execute() error {
    return rootCmd.Execute()
}
