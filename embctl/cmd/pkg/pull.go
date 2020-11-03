package cmd

import (
	"github.com/mj37yhyy/embassy/embctl/exec"
	"github.com/spf13/cobra"
)

var (
	pullCmd = &cobra.Command{
		Use:               "pull",
		Short:             "pull wasm file from embassy server.",
		SilenceUsage:      true,
		DisableAutoGenTag: true,
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				cmd.Help()
				return
			}
			err := exec.Pull(args[0])
			if err != nil {
				cmd.PrintErr(err)
			} else {
				cmd.Printf("pull %s successful!", args[0])
			}
		},
	}
)
