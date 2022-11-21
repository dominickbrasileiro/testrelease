package service

import (
	"fmt"

	"github.com/spf13/cobra"
)

var StopCmd = &cobra.Command{
	Use:   "stop",
	Short: "Stop Docker containers",
	Long:  `Stop Docker containers`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("service stop called")
	},
}
