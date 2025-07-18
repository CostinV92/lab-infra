/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"labtool/cmd/bashrc"
	"labtool/cmd/configure"
	"labtool/cmd/service"
	"labtool/cmd/utilcmd"
	"os"

	"github.com/spf13/cobra"
)

var (
	// rootCmd represents the base command when called without any subcommands
	rootCmd = &cobra.Command{
		Use:   "labtool",
		Short: "Tool for managing lab-infra",
		Long: `Wrapper around various ansible playbooks that lets you configure
and run your ansible environment in a user friendly way.`,
		// Uncomment the following line if your bare application
		// has an action associated with it:
		// Run: func(cmd *cobra.Command, args []string) { fmt.Printf("Command run\n") },
		PersistentPostRun: func(cmd *cobra.Command, args []string) {
			utilcmd.RunCmd()
		},
	}
)

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(configure.ConfigureCmd)
	rootCmd.AddCommand(service.ServiceCmd)
	rootCmd.AddCommand(bashrc.BashrcCmd)
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVarP(&configure.CfgFile, "config", "c", configure.CfgFile, "config file to use or write")
}
