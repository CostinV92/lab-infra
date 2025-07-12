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

// serviceCmd represents the service command
var restartCmd = &cobra.Command{
	Use:   "restart",
	Short: "restart a service",
	Long:  "restart a service",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {

		shellCmd := exec.Command("ansible-playbook")

		shellCmd.Args = append(shellCmd.Args, "-i", configure.Cfg.InventoryPath)
		shellCmd.Args = append(shellCmd.Args, "-e", fmt.Sprintf("target_host=%s", Host))
		shellCmd.Args = append(shellCmd.Args, "-e", fmt.Sprintf("service=%s", Service))
		shellCmd.Args = append(shellCmd.Args, configure.Cfg.PlaybookDir+"/restart_service.yaml")

		shellCmd.Stdout = os.Stdout
		shellCmd.Stderr = os.Stderr

		shellCmd.Run()
	},
}

func init() {

}
