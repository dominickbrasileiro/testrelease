package retriever

import (
	"fmt"

	"github.com/spf13/cobra"
)

var RetrieveCmd = &cobra.Command{
	Use:   "retrieve",
	Short: "Retrieve data from datastore",
	Long:  `Retrieve data from datastore`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("retriever retrieve called")
	},
}
