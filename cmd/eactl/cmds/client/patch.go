package client

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {

	// add the default category flags to the patch command
	AddCategoryFlags(PatchCmd)
}

var PatchCmd = &cobra.Command{
	Use:   "patch",
	Short: "Patch audits on the easy-audit server",
	Long:  `Patch audits on the easy-audit server`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return fmt.Errorf("not implemented")
	},
}
