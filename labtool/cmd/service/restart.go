/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package service

import (
	"labtool/cmd/configure"
	"labtool/cmd/utilcmd"

	"github.com/spf13/cobra"
)

// serviceCmd represents the service command
var restartCmd = &cobra.Command{
	Use:   "restart",
	Short: "restart a service",
	Long:  "restart a service",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		utilcmd.SetPlaybook(configure.Cfg.PlaybookDir + "/restart_service.yaml")
	},
}

func init() {

}
