/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"log"

	"github.com/panthercorp/terradep/internal/configs"
	"github.com/panthercorp/terradep/pkg/files"
	"github.com/spf13/cobra"
)

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Manage terradep configuration",
}

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize terradep/dcf configuration",
	Run: func(cmd *cobra.Command, args []string) {
		userHomeDir, err := files.GetUserHomeDir()
		if err != nil {
			log.Fatal(err)
		}
		typeFlag, err := cmd.Flags().GetString("type")
		if err != nil {
			log.Fatal(err)
		}
		pathFlag, err := cmd.Flags().GetString("path")
		if err != nil {
			log.Fatal(err)
		}
		forceFlag, err := cmd.Flags().GetBool("force")
		if err != nil {
			log.Fatal(err)
		}
		switch typeFlag {
		case "terradep":
			var initializePath string
			if pathFlag == "" {
				initializePath = userHomeDir
			} else {
				initializePath = pathFlag
			}
			if err := configs.InitializeTerradepConfig(initializePath, forceFlag); err != nil {
				log.Fatal(err)
			}
		case "dcf":
			if err := configs.InitializeDcfConfig(pathFlag, forceFlag); err != nil {
				log.Fatal(err)
			}
		default:
			log.Fatal("Invalid config type")
		}
	},
}

// repoCmd represents the repo command
var repoCmd = &cobra.Command{
	Use:   "repo",
	Short: "Manage terradep repositories",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("repo called")
	},
}

func init() {
	rootCmd.AddCommand(configCmd)
	configCmd.AddCommand(initCmd)
	configCmd.AddCommand(repoCmd)

	initCmd.Flags().BoolP("force", "f", false, "Force overwrite of existing terradep config")
	initCmd.Flags().StringP("path", "p", "", "Path to terradep/dcf config")
	initCmd.Flags().StringP("type", "t", "terradep", "Type of config - dcf or terradep")
}
