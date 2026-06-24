package cli

import "github.com/spf13/cobra"

type AddFlagFunc func(c *cobra.Command)

// adds flags and then returns the command
func AppendFlagsFn(cmd *cobra.Command, addFlagFuncs ...AddFlagFunc) *cobra.Command {
	for _, addFlagFunc := range addFlagFuncs {
		addFlagFunc(cmd)
	}
	return cmd
}
