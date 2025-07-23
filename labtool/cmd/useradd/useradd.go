/*
Copyright © 2025 Victor-Gabriel Costin <costinv92@gmail.com>
*/
package useradd

import (
	"bytes"
	"labtool/cmd/configure"
	"labtool/cmd/utilcmd"
	"os/exec"
	"strconv"

	"github.com/spf13/cobra"
)

var (
	host string
	user string
	uid  int
	pass string

	UseraddCmd = &cobra.Command{
		Use:   "useradd",
		Short: "Add a new user to target host",
		Long:  "Add a new user to the target host. The user is added to the sudo group. If the user or the UID already exists, does nothing.",
		Run: func(cmd *cobra.Command, args []string) {
			useraddRun(host, user, uid)
		},
	}
)

func init() {
	UseraddCmd.PersistentFlags().StringVarP(&host, "target-host", "t", "", "host to add the user to")
	UseraddCmd.MarkPersistentFlagRequired("target-host")

	UseraddCmd.PersistentFlags().StringVarP(&user, "user", "u", "", "user to be added")
	UseraddCmd.MarkPersistentFlagRequired("user")

	UseraddCmd.PersistentFlags().IntVarP(&uid, "user-id", "i", 0, "user id of the new user")
	UseraddCmd.MarkPersistentFlagRequired("user-id")

	UseraddCmd.PersistentFlags().StringVarP(&pass, "password", "p", "", "plaintext password")
	UseraddCmd.MarkPersistentFlagRequired("password")
}

func useraddRun(host, user string, uid int) {
	utilcmd.InitRunCmd()
	utilcmd.AddExtraVar("target_host", host)
	utilcmd.AddExtraVar("new_user", user)
	utilcmd.AddExtraVar("uid", strconv.Itoa(uid))

	passHash := generatePasswordHash(pass)
	utilcmd.AddExtraVar("password", passHash)

	utilcmd.SetPlaybook(configure.Cfg.PlaybookDir + "/useradd.yaml")
}

func generatePasswordHash(pass string) string {
	hashPass := bytes.Buffer{}

	makePassCmd := exec.Command("mkpasswd")
	makePassCmd.Args = append(makePassCmd.Args, "--method=sha-512")
	makePassCmd.Args = append(makePassCmd.Args, pass)
	makePassCmd.Stdout = &hashPass

	cobra.CheckErr(makePassCmd.Run())

	return hashPass.String()
}
