package cmd

import (
	"testrelease/cmd/retriever"

	"github.com/spf13/cobra"
)

var RetrieverCmd = &cobra.Command{
	Use:   "retriever",
	Short: "Call retriever methods via gRPC",
	Long:  `Call retriever methods via gRPC`,
}

func init() {
	RetrieverCmd.AddCommand(retriever.RetrieveCmd)
}
