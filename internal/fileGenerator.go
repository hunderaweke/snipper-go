package internal

import (
	"fmt"
	"os"
)

func generateFiles(projectDir string, name string) {
	os.Mkdir(name, 0755)
	os.Chdir(name)
	switch name {
	case "models":
		writeModels("models.go")
	case "routes":
		writeRouters(name + ".go")
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
func writeMain() {
	file, err := os.Create("main.go")
	if err != nil {
		panic(err)
	}
	content := `package main
import (
	"github.com/gofiber/fiber/v2"
)
func main(){
	app := fiber.NewApp()
	app.Get("/", func (c *fiber.Ctx) error {
        return c.SendString("Hello, World!")
    })

    log.Fatal(app.Listen(":3000"))
}
	`
	file.WriteString(content)
	defer file.Close()
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

func writeRouters(fileName string) {
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

func writeEnv() {
	content :=
		`DB_USERNAME=username
DB_PASSWORD=password
DB_NAME=db
DB_PORT=port
DB_HOST=localhost
`
	file, err := os.Create(".env.sample")
	if err != nil {
		panic(err)
	}
	_, err = file.WriteString(content)
	if err != nil {
		panic(err)
	}
}
