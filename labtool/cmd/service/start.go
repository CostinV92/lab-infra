/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package service

import (
	"labtool/cmd/configure"

	"github.com/spf13/cobra"
)

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "start a service",
	Long:  "start a service",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		shellCmd.Args = append(shellCmd.Args, extraVars...)
		shellCmd.Args = append(shellCmd.Args, configure.Cfg.PlaybookDir+"/start_service.yaml")

		shellCmd.Run()
	},
}

func init() {

}
