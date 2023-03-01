package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
)

func main() {

	cmd := "open"
	currentDir, err := os.Getwd()
	log.Println(currentDir)
	if err != nil {
		fmt.Println(err)
	}

	downloadDir := filepath.Join(currentDir, "go.mod")
	log.Println("GOOS : ", runtime.GOOS)

	if runtime.GOOS == "windows" {
		cmd = "explorer"
	} else if runtime.GOOS == "darwin" {
		cmd = "open"
	} else {
		log.Println("Not supported OS")
	}
	// Run wait until it's complete
	if err := exec.Command(cmd, downloadDir).Run(); err != nil {
		log.Fatal("Failed to open downloaded test file!")
	}
}
