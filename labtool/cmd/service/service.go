/*
Copyright © 2025 Victor-Gabriel Costin <costinv92@gmail.com>
*/
package service

import (
	"fmt"
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
			utilcmd.InitRunCmd()
			servicePreRun(service, configure.Cfg.LabEnvDir, configure.Cfg.ServicesDir, configure.Cfg.ServicesEnvDir)
		},
	}
)

func init() {
	ServiceCmd.AddGroup(&configure.CommandsGroup)

	ServiceCmd.AddCommand(startCmd)
	startCmd.GroupID = configure.CommandsGroup.ID

	ServiceCmd.AddCommand(restartCmd)
	restartCmd.GroupID = configure.CommandsGroup.ID

	ServiceCmd.PersistentFlags().StringVarP(&configure.Host, "target-host", "t", "", "host to run the command on")
	ServiceCmd.MarkPersistentFlagRequired("target-host")
	ServiceCmd.PersistentFlags().StringVarP(&service, "service", "s", "", "service to run command for")
	ServiceCmd.MarkPersistentFlagRequired("service")
}

func servicePreRun(service, envDir, servicesDir, servicesEnvDir string) {
	// Add extra-vars
	fmt.Println(servicesEnvDir)
	utilcmd.AddExtraVar("env_dir", envDir)
	utilcmd.AddExtraVar("service", service)
	utilcmd.AddExtraVar("services_dir", servicesDir)
	utilcmd.AddExtraVar("services_env_dir", servicesEnvDir)
}
