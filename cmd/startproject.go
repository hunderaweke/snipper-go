/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/hunderaweke/snipper-go/internal"
	"github.com/spf13/cobra"
)

var startprojectCmd = &cobra.Command{
	Use:   "startproject",
	Short: "A Project starting boilerplate",
	Long:  `A project starting command for creating a go project inside the current directory`,
	Run: func(cmd *cobra.Command, args []string) {
		var projectName string
		if len(args) < 1 {
			fmt.Println("Please provide a project name.")
			return
		}
		projectName = args[0]
		internal.GenerateDirectories(projectName)
	},
}

func init() {
	rootCmd.AddCommand(startprojectCmd)
}
