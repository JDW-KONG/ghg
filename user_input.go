package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("\nSimple Shell")

	fmt.Println("-------------------------------")

	// for loop without parameters acts as a while loop
	for {
		// prints prompt
		fmt.Print("-> ")

		// assign text to equal everything including newline
		text, _ := reader.ReadString('\n')

		// removes trailing newline character
		text = strings.Replace(text, "\n", "", -1)

		if strings.Compare("hi", text) == 0 {
			fmt.Println("hello, Yourself")
		} else if text == "Justin" {
			fmt.Println("Hello Justin! You will win!")
		}
	}
}
