/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/hunderaweke/snipper-go/internal"
	"github.com/spf13/cobra"
)

// addDockerCmd represents the addDocker command
var addDockerCmd = &cobra.Command{
	Use:   "add-docker",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		internal.DockerWriter()
		internal.DockerComposeWriter()
	},
}

func init() {
	rootCmd.AddCommand(addDockerCmd)
}
