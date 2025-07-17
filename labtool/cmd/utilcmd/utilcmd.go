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
	playbook  string
)

func InitCmd() {
	configure.ReadConfigFile()
	os.Setenv("ANSIBLE_CONFIG", configure.Cfg.AnsibleConfigFile)

	// Prepare shellCmd
	shellCmd = exec.Command("ansible-playbook")
	shellCmd.Stdout = os.Stdout
	shellCmd.Stderr = os.Stderr
	shellCmd.Args = append(shellCmd.Args, "-i", configure.Cfg.InventoryPath)

	// Add extra-vars
	extraVars = append(extraVars, "-e", fmt.Sprintf("target_host=%s", configure.Host))
	extraVars = append(extraVars, "-e", fmt.Sprintf("admin_user=%s", configure.Cfg.AdminUser))
}

func AddExtraVar(key, value string) {
	extraVars = append(extraVars, "-e", fmt.Sprintf("%s=%s", key, value))
}

func SetPlaybook(play string) {
	playbook = play
}

func RunCmd() {
	if shellCmd == nil {
		return
	}

	shellCmd.Args = append(shellCmd.Args, extraVars...)
	shellCmd.Args = append(shellCmd.Args, playbook)

	fmt.Println(shellCmd)

	shellCmd.Run()
}
