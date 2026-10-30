/*
Copyright © 2026 Victor-Gabriel Costin <costinv92@gmail.com>
*/
package generateEnv

import (
	"fmt"
	"labtool/cmd/configure"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

// generateEnvCmd represents the generateEnv command
var (
	force          bool
	service        string
	envExamplePath string
	envFilePath    string

	GenerateEnvCmd = &cobra.Command{
		Use:   "generate-env",
		Short: "generate service .env",
		Long:  "Generate a service .env file from its env.local.example",
		Args:  cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			runGenerateEnv()
		},
	}
)

const ErrorServiceDirNoExistsFmt = "directory for service '%s' doesn't exist"
const ErrorServiceEnvDirNoExistsFmt = "env directory for service '%s' doesn't exist"
const ErrorExampleEnvFileNoExistsFmt = "example env file ('env.local.example') for service '%s' doesn't exist"

func validatePaths() error {
	if configure.DryRun {
		return nil
	}

	serviceDir := configure.Cfg.ServicesDir + "/" + service
	serviceEnvDir := configure.Cfg.ServicesEnvDir + "/" + service
	envExamplePath = serviceDir + "/env.local.example"
	envFilePath = serviceEnvDir + "/env.local"

	_, err := os.Stat(serviceDir)
	if err != nil {
		return fmt.Errorf(ErrorServiceDirNoExistsFmt, service)
	}

	_, err = os.Stat(serviceEnvDir)
	if err != nil {
		return fmt.Errorf(ErrorServiceEnvDirNoExistsFmt, service)
	}

	_, err = os.Stat(envExamplePath)
	if err != nil {
		return fmt.Errorf(ErrorExampleEnvFileNoExistsFmt, service)
	}

	return nil
}

func runGenerateEnv() {
	cobra.CheckErr(validatePaths())

	copyEnvExample()
	editEnvFile()
}

func copyEnvExample() {
	_, err := os.Stat(envExamplePath)
	if err == nil && !force {
		return
	}

	cmd := exec.Command("cp")
	cmd.Args = append(cmd.Args, envExamplePath)
	cmd.Args = append(cmd.Args, envFilePath)
	err = cmd.Run()
	fmt.Println(cmd)
}

func editEnvFile() {
	editor := os.Getenv("EDITOR")
	cmd := exec.Command(editor, envFilePath)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	err := cmd.Run()
	if err != nil {
		fmt.Fprint(os.Stderr, "Could not open editor")
		os.Exit(1)
	}
}

func init() {
	//editor := os.Getenv("EDITOR")
	//cmd := exec.Command(editor, "test.txt")
	//cmd.Stdin = os.Stdin
	//cmd.Stdout = os.Stdout
	//err := cmd.Run()
	//if err != nil {
	//	fmt.Fprint(os.Stderr, "Could not open editor")
	//		os.Exit(1)
	//	}

	configure.ReadConfigFile()

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// generateEnvCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// generateEnvCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	GenerateEnvCmd.PersistentFlags().StringVarP(&service, "service", "s", "", "service to generate the env for")
	GenerateEnvCmd.MarkPersistentFlagRequired("service")

	GenerateEnvCmd.PersistentFlags().BoolVarP(&force, "force", "f", false, "overwite the current env file if it exists")
}
