/*
Copyright © 2025 Victor-Gabriel Costin <costinv92@gmail.com
*/
package upgrade

import (
	"labtool/cmd/configure"
	"labtool/cmd/utilcmd"

	"github.com/spf13/cobra"
)

// updateCmd represents the update command
var UpgradeCmd = &cobra.Command{
	Use:   "upgrade",
	Short: "Upgrade a host",
	Long:  `run 'apt update && apt upgrade' on selected host`,
	Run: func(cmd *cobra.Command, args []string) {
		utilcmd.InitRunCmd()
		utilcmd.AddExtraVar("target_host", configure.Host)
		utilcmd.SetPlaybook(configure.Cfg.PlaybookDir + "/upgrade_host.yaml")
	},
}

func init() {
	UpgradeCmd.PersistentFlags().StringVarP(&configure.Host, "target-host", "t", "", "host to run the command on")
	UpgradeCmd.MarkPersistentFlagRequired("target-host")
}
