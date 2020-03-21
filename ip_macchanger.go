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

	fmt.Println("*** Current Settings ***\n")
	err := ExecuteSystemCommand("ip", []string{"link", "show"})
	if err != nil {
		fmt.Println("[-] Unable to print interfaces.\n")
	}

	fmt.Println("\n\nPlease enter the interface: ")
	reader := bufio.NewReader(os.Stdin)

	iface, _ := reader.ReadString('\n')
	iface = strings.Replace(iface, "\n", "", -1)

	fmt.Println("\nYou entered: " + iface + "\n")

	fmt.Println("Please enter the new MAC address: ")
	newMAC, _ := reader.ReadString('\n')
	newMAC = strings.Replace(newMAC, "\n", "", -1)

	fmt.Println("\nYou entered: " + newMAC + "\n")

	fmt.Printf("[+] Updating MAC address on %s to %s\n\n", iface, newMAC)

	fmt.Println("\n[-] Previous Settings\n")
	err = ExecuteSystemCommand("ip", []string{"link", "show", iface})
	if err != nil {
		fmt.Println("[-] Unable to print interface info.\n")
	}

	err = ExecuteSystemCommand("ip", []string{"link", "set", "dev", iface, "address", newMAC})
	if err != nil {
		fmt.Println("[-] Unable to shut down interface.\n")
		fmt.Println("[-] Unable to change MAC address.\n")
	}

	err = ExecuteSystemCommand("ip", []string{"link", "set", "dev", iface, "up"})
	if err != nil {
		fmt.Println("[-] Unable to restore interface.\n")
	}

	fmt.Println("\n\n[+] Updated Settings\n")
	err = ExecuteSystemCommand("ip", []string{"link", "show", iface})
	if err != nil {
		fmt.Println("[-] Unable to print updated info.\n")
	}

	fmt.Println("\n\n")
}
