/*
Copyright © 2025 Victor-Gabriel Costin <costinv92@gmail.com>
*/
package utilcmd

import (
	"fmt"
	"labtool/cmd/configure"
	"os"
	"os/exec"
)

var (
	shellCmd  *exec.Cmd
	extraVars []string
	flags     []string
	playbook  string
)

func BuildCmd() *exec.Cmd {
	if shellCmd == nil {
		return shellCmd
	}

	shellCmd.Args = append(shellCmd.Args, extraVars...)
	shellCmd.Args = append(shellCmd.Args, flags...)
	shellCmd.Args = append(shellCmd.Args, playbook)

	if configure.Verbose {
		shellCmd.Args = append(shellCmd.Args, "-vvv")
	}

	return shellCmd
}

func InitRunCmd() {
	configure.ReadConfigFile()

	// Prepare shellCmd
	shellCmd = exec.Command("ansible-playbook")
	shellCmd.Stdout = os.Stdout
	shellCmd.Stderr = os.Stderr
	shellCmd.Args = append(shellCmd.Args, "-i", configure.Cfg.InventoryPath)

	os.Setenv("ANSIBLE_CONFIG", configure.Cfg.AnsibleConfigFile)

	// Add extra-vars
	extraVars = append(extraVars, "-e", fmt.Sprintf("target_host=%s", configure.Host))
	extraVars = append(extraVars, "-e", fmt.Sprintf("admin_user=%s", configure.Cfg.AdminUser))
}

func AddFlag(flag, value string) {
	flags = append(flags, fmt.Sprintf("-%s=%s", flag, value))
}

func AddExtraVar(key, value string) {
	extraVars = append(extraVars, "-e", fmt.Sprintf("%s=%s", key, value))
}

func SetPlaybook(play string) {
	playbook = play
}

func RunCmd() {
	BuildCmd()

	if !configure.DryRun {
		shellCmd.Run()
	} else {
		fmt.Println(shellCmd)
	}
}
