/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package service

import (
	"fmt"
	"labtool/cmd/configure"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "start a service",
	Long:  "start a service",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		shellCmd := exec.Command("ansible-playbook")

		shellCmd.Args = append(shellCmd.Args, "-i", configure.Cfg.InventoryPath)
		shellCmd.Args = append(shellCmd.Args, "-e", fmt.Sprintf("target_host=%s", Host))
		shellCmd.Args = append(shellCmd.Args, "-e", fmt.Sprintf("service=%s", Service))
		shellCmd.Args = append(shellCmd.Args, "-e", fmt.Sprintf("admin_user=%s", configure.Cfg.AdminUser))
		shellCmd.Args = append(shellCmd.Args, "-e", fmt.Sprintf("scripts_dir=%s", configure.Cfg.ScriptsDir))
		shellCmd.Args = append(shellCmd.Args, "-e", fmt.Sprintf("services_dir=%s", configure.Cfg.ServicesDir))
		shellCmd.Args = append(shellCmd.Args, "-e", fmt.Sprintf("env_dir=%s", configure.Cfg.ServicesEnvDir))
		shellCmd.Args = append(shellCmd.Args, configure.Cfg.PlaybookDir+"/start_service.yaml")

		shellCmd.Stdout = os.Stdout
		shellCmd.Stderr = os.Stderr

		shellCmd.Run()
	},
}

func init() {

}
