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

var (
	Service string
	Host    string

	shellCmd  *exec.Cmd
	extraVars []string
)

// serviceCmd represents the service command
var ServiceCmd = &cobra.Command{
	Use:   "service",
	Short: "service related commands",
	Long:  `Control your services`,
	// Run: func(cmd *cobra.Command, args []string) { },
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		configure.ReadConfigFile()

		os.Setenv("ANSIBLE_CONFIG", configure.Cfg.AnsibleConfigFile)

		// Prepare shellCmd
		shellCmd = exec.Command("ansible-playbook")
		shellCmd.Stdout = os.Stdout
		shellCmd.Stderr = os.Stderr
		shellCmd.Args = append(shellCmd.Args, "-i", configure.Cfg.InventoryPath)

		// Add extra-vars
		extraVars = append(extraVars, "-e", fmt.Sprintf("target_host=%s", Host))
		extraVars = append(extraVars, "-e", fmt.Sprintf("service=%s", Service))
		extraVars = append(extraVars, "-e", fmt.Sprintf("admin_user=%s", configure.Cfg.AdminUser))
		// TODO: do i need scripts_dir here?
		extraVars = append(extraVars, "-e", fmt.Sprintf("scripts_dir=%s", configure.Cfg.ScriptsDir))
		extraVars = append(extraVars, "-e", fmt.Sprintf("services_dir=%s", configure.Cfg.ServicesDir))
		extraVars = append(extraVars, "-e", fmt.Sprintf("env_dir=%s", configure.Cfg.ServicesEnvDir))
	},
}

func init() {
	ServiceCmd.AddCommand(startCmd)
	ServiceCmd.AddCommand(restartCmd)

	ServiceCmd.PersistentFlags().StringVarP(&Service, "service", "s", "", "service to run command for")
	ServiceCmd.PersistentFlags().StringVarP(&Host, "target-host", "t", "", "host to run the command on")
	ServiceCmd.MarkPersistentFlagRequired("service")
	ServiceCmd.MarkPersistentFlagRequired("target-host")
}
