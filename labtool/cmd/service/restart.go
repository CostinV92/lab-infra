/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package service

import (
	"labtool/cmd/configure"

	"github.com/spf13/cobra"
)

// serviceCmd represents the service command
var restartCmd = &cobra.Command{
	Use:   "restart",
	Short: "restart a service",
	Long:  "restart a service",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		shellCmd.Args = append(shellCmd.Args, extraVars...)
		shellCmd.Args = append(shellCmd.Args, configure.Cfg.PlaybookDir+"/restart_service.yaml")

		shellCmd.Run()
	},
}

func init() {

}
