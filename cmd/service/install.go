package service

import (
	"fmt"

	"github.com/spf13/cobra"
)

var InstallCmd = &cobra.Command{
	Use:   "install",
	Short: "Install Testrelease components using Docker and Docker Compose",
	Long:  `Install Testrelease components using Docker and Docker Compose`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("service install called")
	},
}
