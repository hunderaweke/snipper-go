package internal

import (
	"fmt"
	"os"
)

func GenerateDirectories(projectName string) {
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
	github.com/joho/godotenv
)
	`)
	defer file.Close()
	generateFiles(projectName, "models")
	generateFiles(projectName, "handlers")
	generateFiles(projectName, "config")
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
