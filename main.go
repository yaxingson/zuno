package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"zuno/command"
)

var rootCmd = &cobra.Command{
	Use:   "zuno",
	Short: "Zuno is an AI-based all-rounder",
	Long:  "An intelligent assistant built on a series of models such as gpt, gemini and grok",
	Run: func(cmd *cobra.Command, args []string) {
		for {
			break
		}
	},
}

func main() {
	rootCmd.AddCommand(command.ProviderCmd)
	rootCmd.AddCommand(command.UiCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
