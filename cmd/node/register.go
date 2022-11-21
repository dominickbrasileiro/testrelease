package node

import (
	"fmt"

	"github.com/spf13/cobra"
)

var RegisterCmd = &cobra.Command{
	Use:   "register",
	Short: "Register a Testrelease node",
	Long:  `Register a Testrelease node`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("node register called")
	},
}
