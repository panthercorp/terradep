/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// dependencyCmd represents the dependency command
var dependencyCmd = &cobra.Command{
	Use:   "dependency",
	Short: "Manage terraform dependencies",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("dependency called")
	},
}

// installCmd represents the install command
var installCmd = &cobra.Command{
	Use:   "install",
	Short: "Install all dependencies present in dcf.yml",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("install called")
	},
}

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update all dependencies present in dcf.yml",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("update called")
	},
}

// removeCmd represents the remove command
var removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove all downloaded dependencies",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("remove called")
	},
}

// buildCmd represents the build command
var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "Build all dependencies present in dcf.yml and upload to a remote repository",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("build called")
	},
}

func init() {
	rootCmd.AddCommand(dependencyCmd)
	dependencyCmd.AddCommand(installCmd)
	dependencyCmd.AddCommand(updateCmd)
	dependencyCmd.AddCommand(removeCmd)
	dependencyCmd.AddCommand(buildCmd)
}
