package cmd

import (
	"github.com/Nuyoahch/conflux/ipconf"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(ipConfCmd)
}

var ipConfCmd = &cobra.Command{
	Use: "ipconf",
	Run: IpConfHandle,
}

func IpConfHandle(cmd *cobra.Command, args []string) {
	ipconf.RunMain()
}
