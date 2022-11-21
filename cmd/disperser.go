package cmd

import (
	"testrelease/cmd/disperser"

	"github.com/spf13/cobra"
)

var DisperserCmd = &cobra.Command{
	Use:   "disperser",
	Short: "Call disperser methods via gRPC",
	Long:  `Call disperser methods via gRPC`,
}

func init() {
	DisperserCmd.AddCommand(disperser.DisperseCmd)
	DisperserCmd.AddCommand(disperser.EncodeCmd)
}
