package main

import (
	"os"

	"github.com/jnnkrdb/easy-audit/cmd/eactl/cmds"
)

func main() {
	err := cmds.RootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
