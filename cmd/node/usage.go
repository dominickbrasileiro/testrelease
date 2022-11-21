package node

import (
	"fmt"

	"github.com/spf13/cobra"
)

var UsageCmd = &cobra.Command{
	Use:   "usage",
	Short: "Check a Testrelease node disk usage",
	Long:  `Check a Testrelease node disk usage`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("node usage called")
	},
}
