package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

func main() {

	fmt.Println("Welcome to Betsys app, enter your command...")

	command := flag.String("command", "no-command", "command to run")

	for {
		runCommand(*command)

		scanner := bufio.NewScanner(os.Stdin)
		fmt.Println(`
________________________________________________________
command list:
	- list-clients				
		retrieves a list of unregistered clients	
	- list-registered-clients				
		retrieves a list of registered clients
	- register-client
		registers a new client
	- register-client-to-server
		registers a client to server with userID
________________________________________________________
		`)
		scanner.Scan()
		*command = scanner.Text()

	}
}

func runCommand(command string) {

	// TODO : create each command's function
	switch command {
	case "client-list":
		fmt.Println("input command: client-list")
	case "register-client":
		fmt.Println("input command: register-client")
	case "exit":
		os.Exit(0)
	default:
		fmt.Println("Command is not valid", command)
	}

}
