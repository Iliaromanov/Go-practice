package main

import (
	"github.com/spf13/cobra"
	"os"
	"fmt"
)

func main() {
	var appCmd = &cobra.Command {
		Use:   "app",
		Short: "The AppName CLI",
		Run: func(cmd *cobra.Command, args []string) {
		},
	}

	err := appCmd.Execute()
	if err != nil {
		fmt.Fprintln(os.Stderr, err) // output error into Stderr (instead of stdout)
		os.Exit(1)
	}
}