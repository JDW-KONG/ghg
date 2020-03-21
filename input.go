package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("*** MAC Changer ***")
	fmt.Println("-------------------------------------\n")

	fmt.Println("Please enter the interface: ")
	reader := bufio.NewReader(os.Stdin)

	text, _ := reader.ReadString('\n')
	text = strings.Replace(text, "\n", "", -1)

	fmt.Println("You entered " + text + ".")
}
