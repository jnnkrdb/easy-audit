package post

import (
	"fmt"

	"github.com/spf13/cobra"
)

var PostCmd = &cobra.Command{
	Use:   "post",
	Short: "Post audits to the easy-audit server",
	Long:  `Post audits to the easy-audit server`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return fmt.Errorf("not implemented")
	},
}

func init() {
}
