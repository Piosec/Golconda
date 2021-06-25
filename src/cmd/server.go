package cmd

import (
   "golconda/src/server"
   "golconda/src"
   "github.com/spf13/cobra"
   "errors"
)

var serverCmd = &cobra.Command{
  Use:   "server",
  Short: "Use server to listen reverse ports.",
  Long: `A Fast server reverse port discovery made with love.
         Please use it with love too <3.`,
  PreRunE: func(cmd *cobra.Command, args []string) error{

       if dump {
            if target == "" || src.ValidateIP(target) == false{
                return errors.New("The target must be set or a valid IPv4.")
            }
            if interface_name == "" {
                return  errors.New("The interface name must be set.")
            }
      }
      return nil
    },
    Run: func(cmd *cobra.Command, args []string) {

      if dump {
        server.MonitorInterface(interface_name,target)
      }
      if ports != "" {
        server.PortHandlers(ports)
      }
      if topports >0 && topports < 65535 {
        //src.GenerateTopPorts(10)
        server.PortHandlers(src.GetTopPorts(topports))
      }
  },
}

func init() {
    serverCmd.Flags().StringVarP(&ports,"ports","p","","Listening to the following ports. Ex: 21,22,8080 or 22-8080")
    serverCmd.Flags().StringVarP(&target,"target","t","127.0.0.1","Target to check reverse ports coming from. Ex: 127.0.0.1")
    serverCmd.Flags().StringVarP(&interface_name,"interface","i","","Interface listening to. Ex: lo")
    serverCmd.Flags().BoolVarP(&dump,"dump","d",false,"Monitor incoming connection from remote target.")
    serverCmd.Flags().IntVarP(&topports,"top-ports","",100,"Listening to the most known ports.")
    rootCmd.AddCommand(serverCmd)
}
