package cmd

import (
	"github.com/mj37yhyy/embassy/embctl/exec"
	"github.com/spf13/cobra"
)

var (
	userName string
	password string
	loginCmd = &cobra.Command{
		Use:               "login",
		Short:             "login to embassy server.",
		SilenceUsage:      true,
		DisableAutoGenTag: true,
		Run: func(cmd *cobra.Command, args []string) {
			if len(args[0]) == 0 || len(userName) == 0 || len(password) == 0 {
				cmd.Help()
				return
			}
			err := exec.Login(args[0], userName, password)
			if err != nil {
				cmd.PrintErr(err)
			} else {
				cmd.Println("login successful!")
			}
		},
	}
)

func init() {
	loginCmd.PersistentFlags().StringVarP(&userName, "user", "u", "", "user name for embassy server")
	loginCmd.PersistentFlags().StringVarP(&password, "password", "p", "", "password for embassy server")
}
