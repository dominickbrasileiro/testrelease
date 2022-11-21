package service

import (
	"fmt"

	"github.com/spf13/cobra"
)

var StartCmd = &cobra.Command{
	Use:   "start",
	Short: "Start Docker containers",
	Long:  `Start Docker containers`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("service start called")
	},
}
