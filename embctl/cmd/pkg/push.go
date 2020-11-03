package cmd

import (
	"github.com/mj37yhyy/embassy/embctl/exec"
	"github.com/spf13/cobra"
)

var (
	pushCmd = &cobra.Command{
		Use:               "push",
		Short:             "push wasm file to embassy server.",
		SilenceUsage:      true,
		DisableAutoGenTag: true,
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				cmd.Help()
				return
			}
			err := exec.Push(args[0])
			if err != nil {
				cmd.PrintErr(err)
			} else {
				cmd.Printf("push %s successful!", args[0])
			}
		},
	}
)
