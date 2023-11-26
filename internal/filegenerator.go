package internal

import (
	"fmt"
	"os"
	"os/exec"
)

func GenerateDirectories(projectName string) {
	os.Mkdir(projectName, 0755)
	os.Chdir(projectName)
	cmd := exec.Command("go", "mod", "init", projectName)
	_, err := cmd.Output()
	if err != nil {
		panic(err)
	}
	generateFiles(projectName, "models")
	generateFiles(projectName, "handlers")
	generateFiles(projectName, "config")
	generateFiles(projectName, "controllers")
	generateFiles(projectName, "routes")
	cmd = exec.Command("go", "mod", "tidy")
	_, err = cmd.Output()
	if err != nil {
		panic(err)
	}
}
func generateFiles(projectDir string, name string) {
	os.Mkdir(name, 0755)
	os.Chdir(name)
	switch name {
	case "models":
		writeModels("models.go")
	case "routes":
		writeRoutes(name + ".go")
	case "controllers":
		writeControllers(name + ".go")
	case "handlers":
		writeHandlers(name + ".go")
	case "cmd":
		writeCmd(name + ".go")
	case "config":
		writeConfig("database.go")
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
func writeCmd(fileName string) {
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println("error ", err)
		return
	}
	content := `package main 
// This is the place for creating your handlers
func main(){
	// A place for running the code completely

}
`
	_, err = file.WriteString(content)
	if err != nil {
		panic(err)
	}
}

func writeRoutes(fileName string) {
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println("error ", err)
		return
	}
	content := `package routes 
// This is the place for creating your routes

`
	_, err = file.WriteString(content)
	if err != nil {
		panic(err)
	}
}
func writeControllers(fileName string) {
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println("error ", err)
		return
	}
	content := `package controllers 
// This is the place for creating your controllers
func init(){
	// Some of the startup codes here
}

`
	_, err = file.WriteString(content)
	if err != nil {
		panic(err)
	}
}
func writeConfig(fileName string) {
	content := `package config
import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"github.com/joho/godotenv"
)

var db *gorm.DB

func Connect() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading the .env file")
	}
	dbUserName := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")
	dbHost := os.Getenv("DB_HOST")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		dbUserName,
		dbPassword,
		dbHost,
		dbPort,
		dbName,
	)
	d, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	db = d
}
func GetDB() *gorm.DB {
	return db
}`
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println("error ", err)
		return
	}
	_, err = file.WriteString(content)
	if err != nil {
		panic(err)
	}
}
