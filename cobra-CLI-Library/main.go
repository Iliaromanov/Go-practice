package main

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use: "example",
		Short: "An example cobra program",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Hello from root command")
		},
	}
)

func init() {
	
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}