/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "terradep",
	Short: "Manage terraform module dependencies with terradep easily",
	Long: `There is no version control support for modules present natively in terraform.
With terradep, you can version control all of your modules and manage them as well.
You can use modules published by others on the default public repository.
This practice avoids writing code repeatedly and enables code reuse.`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
