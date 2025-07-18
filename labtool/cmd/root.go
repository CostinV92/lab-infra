/*
Copyright © 2025 Victor-Gabriel Costin <costinv92@gmail.com>
*/
package cmd

import (
	"labtool/cmd/bashrc"
	"labtool/cmd/configure"
	"labtool/cmd/service"
	"labtool/cmd/upgrade"
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
	rootCmd.AddGroup(&configure.CommandsGroup)
	rootCmd.AddGroup(&configure.HelpGroup)

	rootCmd.SetHelpCommandGroupID(configure.HelpGroup.ID)
	rootCmd.SetCompletionCommandGroupID(configure.HelpGroup.ID)

	rootCmd.AddCommand(configure.ConfigureCmd)
	configure.ConfigureCmd.GroupID = configure.HelpGroup.ID

	rootCmd.AddCommand(service.ServiceCmd)
	service.ServiceCmd.GroupID = configure.CommandsGroup.ID

	rootCmd.AddCommand(bashrc.BashrcCmd)
	bashrc.BashrcCmd.GroupID = configure.CommandsGroup.ID

	rootCmd.AddCommand(upgrade.UpgradeCmd)
	upgrade.UpgradeCmd.GroupID = configure.CommandsGroup.ID

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	rootCmd.PersistentFlags().StringVarP(&configure.CfgFile, "config", "c", configure.CfgFile, "config file to use or write")
}
