/*
Copyright © 2025 Victor-Gabriel Costin <costinv92@gmail.com>
*/
package provision

import (
	"labtool/cmd/configure"
	"labtool/cmd/utilcmd"

	"github.com/spf13/cobra"
)

// provisionCmd represents the provision command
var ProvisionCmd = &cobra.Command{
	Use:   "provision",
	Short: "Provision lab env on a server",
	Long:  `Copy env files needed to run services to a server`,
	Run: func(cmd *cobra.Command, args []string) {
		utilcmd.InitRunCmd()
		utilcmd.AddExtraVar("env_dir", configure.Cfg.LabEnvDir)
		utilcmd.AddExtraVar("scripts_dir", configure.Cfg.ScriptsDir)
		utilcmd.SetPlaybook(configure.Cfg.PlaybookDir + "/provision.yaml")
	},
}

func init() {
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// provisionCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// provisionCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	ProvisionCmd.PersistentFlags().StringVarP(&configure.Host, "target-host", "t", "", "host to provision the env on")
	ProvisionCmd.MarkPersistentFlagRequired("target-host")
}
