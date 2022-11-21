package node

import (
	"fmt"

	"github.com/spf13/cobra"
)

var DeregisterCmd = &cobra.Command{
	Use:   "deregister",
	Short: "Deregister a Testrelease node",
	Long:  `Deregister a Testrelease node`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("node deregister called")
	},
}
