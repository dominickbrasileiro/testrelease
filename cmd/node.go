package cmd

import (
	"testrelease/cmd/node"

	"github.com/spf13/cobra"
)

var NodeCmd = &cobra.Command{
	Use:   "node",
	Short: "Call node methods via gRPC",
	Long:  `Call node methods via gRPC`,
}

func init() {
	NodeCmd.AddCommand(node.DeregisterCmd)
	NodeCmd.AddCommand(node.RegisterCmd)
	NodeCmd.AddCommand(node.UsageCmd)
}
