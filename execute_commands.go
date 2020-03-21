package main

import (
	"log"
	"os"
	"os/exec"
)

func main() {
	cmd := exec.Command("ls")

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()

	if err != nil {
		log.Fatal(err)
	}
}
