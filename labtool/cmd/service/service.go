/*
Copyright © 2025 Victor-Gabriel Costin <costinv92@gmail.com>
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
			servicePreRun(service, configure.Cfg.ServicesDir, configure.Cfg.ServicesEnvDir)
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

func servicePreRun(service, servicesDir, servicesEnvDir string) {
	utilcmd.InitRunCmd()

	// Add extra-vars
	utilcmd.AddExtraVar("service", service)
	utilcmd.AddExtraVar("services_dir", servicesDir)
	utilcmd.AddExtraVar("env_dir", servicesEnvDir)
}
