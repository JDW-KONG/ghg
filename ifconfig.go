package main

import (
	"fmt"
	"os/exec"
)

func main() {
	cmd := exec.Command("ifconfig")

	output, err := cmd.Output()

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(output))
}
