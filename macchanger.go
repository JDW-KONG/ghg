package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

// ExecuteSystemCommand executes system commands
// accepts command and arguments
func ExecuteSystemCommand(command string, arguments []string) (err error) {
	args := arguments
	cmd := exec.Command(command, args...)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	err = cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	return
}

func main() {
	fmt.Println("\n*** MAC Changer ***")
	fmt.Println("-------------------------------------\n")

	fmt.Println("Please enter the interface: ")
	reader := bufio.NewReader(os.Stdin)

	iface, _ := reader.ReadString('\n')
	iface = strings.Replace(iface, "\n", "", -1)

	fmt.Println("\nYou entered: " + iface + "\n")

	fmt.Println("Please enter the new MAC address: ")
	newMAC, _ := reader.ReadString('\n')
	newMAC = strings.Replace(newMAC, "\n", "", -1)

	fmt.Println("\nYou entered: " + newMAC + "\n")

	err := ExecuteSystemCommand("ifconfig", []string{iface})
	if err != nil {
		fmt.Println("[-] Unable to print interface info.\n")
	}

	err = ExecuteSystemCommand("ifconfig", []string{iface, "down"})
	if err != nil {
		fmt.Println("[-] Unable to shut down interface.\n")
	}

	err = ExecuteSystemCommand("ifconfig", []string{iface, "hw", "ether", newMAC})
	if err != nil {
		fmt.Println("[-] Unable to change MAC address.\n")
	}

	err = ExecuteSystemCommand("ifconfig", []string{iface, "up"})
	if err != nil {
		fmt.Println("[-] Unable to restore interface.\n")
	}

}
