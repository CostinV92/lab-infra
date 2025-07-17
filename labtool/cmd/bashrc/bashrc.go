/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package bashrc

import (
	"labtool/cmd/configure"
	"labtool/cmd/utilcmd"

	"github.com/spf13/cobra"
)

// bashrcCmd represents the bashrc command
var BashrcCmd = &cobra.Command{
	Use:   "bashrc",
	Short: "Provision bashrc on a server",
	Long:  `Combine a global and a local (per host) bash rc and provision it on the target host`,
	Run: func(cmd *cobra.Command, args []string) {
		utilcmd.InitCmd()
		utilcmd.AddExtraVar("scripts_dir", configure.Cfg.ScriptsDir)
		utilcmd.SetPlaybook(configure.Cfg.PlaybookDir + "/provision_bashrc.yaml")
	},
}

func init() {
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// bashrcCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// bashrcCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
