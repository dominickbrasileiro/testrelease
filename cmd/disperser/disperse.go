package disperser

import (
	"fmt"

	"github.com/spf13/cobra"
)

var DisperseCmd = &cobra.Command{
	Use:   "disperse",
	Short: "Disperse encoded data to a new datastore",
	Long:  `Disperse encoded data to a new datastore`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("disperser disperse called")
	},
}
