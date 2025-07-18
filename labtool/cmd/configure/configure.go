/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package configure

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type config struct {
	AdminUser         string
	PlaybookDir       string
	InventoryPath     string
	AnsibleConfigFile string
	ServicesDir       string
	ServicesEnvDir    string
	ScriptsDir        string
}

const ErrorConfDirNoExistsFmt = "directory '%s' doesn't exist"
const ErrorConfFileNoExistsFmt = "file '%s' doesn't exist"
const ErrorConfPathNoDirFmt = "%s must be directory"
const ErrorConfPathNoFileFmt = "%s must be file"

var (
	CfgFile string
	Cfg     config
	Host    string
	homeDir string

	// configureCmd represents the configure command
	ConfigureCmd = &cobra.Command{
		Use:   "configure",
		Short: "Configure your labtool env",
		Long: `Generate a config file used by labtool to find and run the ansible scripts.
Warning this will overwrite any existing configuration file`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Warning this will overwrite the provided config file if it exists")

			readConfigFromInput()
			saveConfigToFile()
		},
	}
)

func init() {
	getHomeDir()
}

func SetConfigFile() {
	if CfgFile == "" {
		CfgFile = homeDir + "/.labtool.env"
	}

	viper.SetConfigType("env")
	viper.SetConfigFile(CfgFile)
	fmt.Println("Using config file:", viper.ConfigFileUsed())
}

func ReadConfigFile() {
	if err := viper.ReadInConfig(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: config file not found. Run 'labtool configure' to create it\n")
		os.Exit(1)
	}

	Cfg.AdminUser = viper.GetString("admin_user")
	Cfg.InventoryPath = viper.GetString("ansible_inventory_path")
	Cfg.PlaybookDir = viper.GetString("ansible_playbook_dir")
	Cfg.AnsibleConfigFile = viper.GetString("ansible_config_file")
	Cfg.ScriptsDir = viper.GetString("scripts_dir")
	Cfg.ServicesDir = viper.GetString("services_dir")
	Cfg.ServicesEnvDir = viper.GetString("services_env_dir")
}

func getHomeDir() {
	// Find home directory.
	var err error
	homeDir, err = os.UserHomeDir()
	cobra.CheckErr(err)
}

func appendHomeDir(path *string) {
	*path = homeDir + "/" + *path
}

func discardStdin() {
	var discard string
	for {
		_, err := fmt.Scanln(&discard)
		if err == nil || err.Error() == "unexpected newline" {
			break
		}
	}
}

func validatePath(name, path string, isDir bool) error {
	info, err := os.Stat(path)
	if err != nil {
		if isDir {
			return fmt.Errorf(ErrorConfDirNoExistsFmt, path)
		}
		return fmt.Errorf(ErrorConfFileNoExistsFmt, path)
	}

	if isDir && !info.IsDir() {
		return fmt.Errorf(ErrorConfPathNoDirFmt, name)
	}

	if !isDir && info.IsDir() {
		return fmt.Errorf(ErrorConfPathNoFileFmt, name)
	}

	return nil
}

func readStringToVar(msg string, varString *string) {
	fmt.Printf("%s: ", msg)
	if _, err := fmt.Scanln(varString); err != nil {
		fmt.Println("Error: input must not contain spaces.")
		discardStdin()
		os.Exit(1)
	}
}

func readAdminUser() {
	readStringToVar("Admin user", &Cfg.AdminUser)
}

func readPlaybookDir() {
	readStringToVar("Ansible playbook directory (relatinve to $HOME)", &Cfg.PlaybookDir)
	appendHomeDir(&Cfg.PlaybookDir)
	cobra.CheckErr(validatePath("playbook directory", Cfg.PlaybookDir, true))
}

func readInventoryPath() {
	readStringToVar("Ansible inventory path (relative to $HOME)", &Cfg.InventoryPath)
	appendHomeDir(&Cfg.InventoryPath)
	cobra.CheckErr(validatePath("inventory path", Cfg.InventoryPath, false))
}

func readAnsibleConfigFile() {
	readStringToVar("Ansible config file (relative to $HOME)", &Cfg.AnsibleConfigFile)
	appendHomeDir(&Cfg.AnsibleConfigFile)
	cobra.CheckErr(validatePath("ansible config path", Cfg.AnsibleConfigFile, false))
}

func readServiceDir() {
	readStringToVar("Services directory (relative to $HOME)", &Cfg.ServicesDir)
	appendHomeDir(&Cfg.ServicesDir)
	cobra.CheckErr(validatePath("services directory", Cfg.ServicesDir, true))
}

func readServiceEnvDir() {
	readStringToVar("Services environement directory (relative to $HOME)", &Cfg.ServicesEnvDir)
	appendHomeDir(&Cfg.ServicesEnvDir)
	cobra.CheckErr(validatePath("services env directory", Cfg.ServicesDir, true))
}

func readScriptsDir() {
	readStringToVar("Scripts directory (relative to $HOME)", &Cfg.ScriptsDir)
	appendHomeDir(&Cfg.ScriptsDir)
	cobra.CheckErr(validatePath("scripts directory", Cfg.ScriptsDir, true))
}

func readConfigFromInput() {
	readAdminUser()
	readPlaybookDir()
	readInventoryPath()
	readAnsibleConfigFile()
	readServiceDir()
	readServiceEnvDir()
	readScriptsDir()
}

func saveConfigToFile() {
	viper.Set("admin_user", Cfg.AdminUser)
	viper.Set("ansible_inventory_path", Cfg.InventoryPath)
	viper.Set("ansible_config_file", Cfg.AnsibleConfigFile)
	viper.Set("ansible_playbook_dir", Cfg.PlaybookDir)
	viper.Set("scripts_dir", Cfg.ScriptsDir)
	viper.Set("services_dir", Cfg.ServicesDir)
	viper.Set("setvices_env_dir", Cfg.ServicesEnvDir)

	err := viper.WriteConfigAs(CfgFile)
	if err != nil {
		cobra.CheckErr(err)
	}
}
