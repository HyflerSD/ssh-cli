package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "hssh",
	Short: "Testing",
	Long:  "Testing long",
	Run: func(command *cobra.Command, args []string) {
		fmt.Println("this shit actually works?")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
