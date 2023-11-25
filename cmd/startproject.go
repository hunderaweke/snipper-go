/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

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
		generateDirectories(projectName)
	},
}

func generateDirectories(projectName string) {
	os.Mkdir(projectName, 0755)
	os.Chdir(projectName)
	file, err := os.Create("go.mod")
	if err != nil {
		panic(err)
	}
	file.WriteString(`module github.com/hunderaweke/snipper-go

go 1.21.4

require (
	github.com/andybalholm/brotli v1.0.5 // indirect
	github.com/gofiber/fiber/v2 v2.51.0 // indirect
	github.com/google/uuid v1.4.0 // indirect
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/klauspost/compress v1.16.7 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/mattn/go-runewidth v0.0.15 // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	github.com/valyala/fasthttp v1.50.0 // indirect
	github.com/valyala/tcplisten v1.0.0 // indirect
	golang.org/x/sys v0.14.0 // 
)
	`)
	defer file.Close()
	generateFiles(projectName, "models")
	generateFiles(projectName, "handlers")
}
func generateFiles(projectDir string, name string) {
	os.Mkdir(name, 0755)
	os.Chdir(name)
	switch name {
	case "models":
		writeModels("models.go")
	case "routers":
		break
	case "controllers":
		break
	case "handlers":
		writeHandlers(name + ".go")
	case "cmd":
		break
	}
	os.Chdir("..")
}
func writeModels(fileName string) {
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println("error ", err)
		return
	}
	content := `package models 
// This is the place for creating your database models
import (
	"gorm.io/gorm"
	"gorm.io/driver/sqlite"
)
`
	_, err = file.WriteString(content)
	if err != nil {
		panic(err)
	}
}
func writeHandlers(fileName string) {
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println("error ", err)
		return
	}
	content := `package handlers 
// This is the place for creating your handlers

`
	_, err = file.WriteString(content)
	if err != nil {
		panic(err)
	}
}
func init() {
	rootCmd.AddCommand(startprojectCmd)
}
