package internal

import (
	"fmt"
	"os"
	"os/exec"
)

func GenerateDirectories(projectName string) {
	os.Mkdir(projectName, 0755)
	os.Chdir(projectName)
	writeMain()
	writeEnv()
	cmd := exec.Command("go", "mod", "init", projectName)
	_, err := cmd.Output()
	if err != nil {
		fmt.Println("The directory already exist...")
		return
	}
	generateFiles(projectName, "models")
	generateFiles(projectName, "handlers")
	generateFiles(projectName, "config")
	generateFiles(projectName, "controllers")
	generateFiles(projectName, "routers")
	cmd = exec.Command("go", "mod", "tidy")
	_, err = cmd.Output()
	if err != nil {
		panic(err)
	}
}
