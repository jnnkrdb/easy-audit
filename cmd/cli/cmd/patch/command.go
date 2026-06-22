package patch

import (
	"fmt"

	"github.com/spf13/cobra"
)

var PatchCmd = &cobra.Command{
	Use:   "patch",
	Short: "Patch audits on the easy-audit server",
	Long:  `Patch audits on the easy-audit server`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return fmt.Errorf("not implemented")
	},
}

func init() {
}
