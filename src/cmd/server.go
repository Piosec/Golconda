package cmd

import (
   "fmt"
   // "os"i
   "golconda/src/server"
   "golconda/src"
   "github.com/spf13/cobra"
)

var serverCmd = &cobra.Command{
  Use:   "server",
  Short: "Use server to listen reverse ports.",
  Long: `A Fast server reverse port discovery made with love.
         Please use it with love too <3.`,
  Run: func(cmd *cobra.Command, args []string) {
      if ports != "" {
        server.PortHandlers(ports)
      }
      if topports >0 && topports < 65535 {
          server.PortHandlers(src.GetTopPorts(topports))
      }else{
        fmt.Println("Top ports must be between 1 and 65535")
      }
  },
}

func init() {
    serverCmd.Flags().StringVarP(&ports,"ports","p","","Listening to the following ports. Ex: 21,22,8080 or 22-8080")
    serverCmd.Flags().IntVarP(&topports,"top-ports","",100,"Listening to the most known ports.")
    rootCmd.AddCommand(serverCmd)
}
