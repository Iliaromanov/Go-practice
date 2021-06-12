package main

import (
	"fmt"
	"log"
	"strings"
	"github.com/spf13/cobra"
)

var (
	localRootFlag bool // Only available for rootCmd
	persistRootFlag bool // Available for all commands

	rootCmd = &cobra.Command{
		Use: "example",
		Short: "An example cobra program",
	}

	echoCmd = &cobra.Command{
		Use: "echo [strings to echo]",
		Short: "prints given strings to stdout",
		Args: cobra.MinimumNArgs(1), // Must give at least one argument
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Echo: " + strings.Join(args, " "))
		},
	}
)

func init() {
	rootCmd.PersistentFlags().BoolVarP(&persistRootFlag, "persistFlag", "p", false, "a persistent root flag")
	rootCmd.Flags().BoolVarP(&localRootFlag, "localFlag", "l", false, "a local root flag")
	rootCmd.AddCommand(echoCmd)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}