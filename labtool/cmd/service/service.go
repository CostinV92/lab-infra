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
var (
	service string

	ServiceCmd = &cobra.Command{
		Use:   "service",
		Short: "service related commands",
		Long:  `Control your services`,
		// Run: func(cmd *cobra.Command, args []string) { },
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			configure.SetConfigFile()
			utilcmd.InitCmd()

			// Add extra-vars
			utilcmd.AddExtraVar("service", service)
			utilcmd.AddExtraVar("services_dir", configure.Cfg.ServicesDir)
			utilcmd.AddExtraVar("env_dir", configure.Cfg.ServicesEnvDir)
		},
	}
)

func init() {
	ServiceCmd.AddCommand(startCmd)
	ServiceCmd.AddCommand(restartCmd)

	ServiceCmd.PersistentFlags().StringVarP(&configure.Host, "target-host", "t", "", "host to provision bashrc on")
	ServiceCmd.MarkPersistentFlagRequired("target-host")
	ServiceCmd.PersistentFlags().StringVarP(&service, "service", "s", "", "service to run command for")
	ServiceCmd.MarkPersistentFlagRequired("service")
}
