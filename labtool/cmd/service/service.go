/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package service

import (
	"labtool/cmd/configure"
	"os"

	"github.com/spf13/cobra"
)

var (
	Service string
	Host    string
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
