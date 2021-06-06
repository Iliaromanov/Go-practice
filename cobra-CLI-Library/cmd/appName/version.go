package main

import (
	"fmt"
	"github.com/spf13/cobra"
)

var versionCmd = & cobra.Command {
	Use:	"version",
	Short:  "Describes version.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Version: 0.1.0 -beta Version Describer")
	},
}