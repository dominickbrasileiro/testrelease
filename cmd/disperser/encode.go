package disperser

import (
	"fmt"

	"github.com/spf13/cobra"
)

var EncodeCmd = &cobra.Command{
	Use:   "encode",
	Short: "Encode data to be dispersed",
	Long:  `Encode data to be dispersed`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("disperser encode called")
	},
}
