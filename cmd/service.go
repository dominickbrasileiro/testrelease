package cmd

import (
	"testrelease/cmd/service"

	"github.com/spf13/cobra"
)

var ServiceCmd = &cobra.Command{
	Use:   "service",
	Short: "Manage Testrelease components",
	Long:  `Manage Testrelease components`,
}

func init() {
	ServiceCmd.AddCommand(service.InstallCmd)
	ServiceCmd.AddCommand(service.StartCmd)
	ServiceCmd.AddCommand(service.StopCmd)
}
