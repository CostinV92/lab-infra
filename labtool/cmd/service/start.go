/*
Copyright © 2025 Victor-Gabriel Costin <costinv92@gmail.com>
*/
package service

import (
	"labtool/cmd/configure"
	"labtool/cmd/utilcmd"

	"github.com/spf13/cobra"
)

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "start a service",
	Long:  "start a service",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		startRun()
	},
}

func init() {

}

func startRun() {
	utilcmd.SetPlaybook(configure.Cfg.PlaybookDir + "/start_service.yaml")
}
