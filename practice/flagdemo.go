package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {

	// Create new FlagSet for "add" subcommand
	addCommand := flag.NewFlagSet("add", flag.ExitOnError)
	// Define a string flag called name
	userName := addCommand.String("name", "", "Name of user")

	// Create new FlagSet for "remove" subcommand
	removeCommand := flag.NewFlagSet("remove", flag.ExitOnError)
	// Define a boolean flag to check that user's home directory should be removed
	removeDirectory := removeCommand.Bool("d", false, "Remove home directory of a user")

	// If no subcommand is entered then we need to exit from application
	if len(os.Args) < 2 {
		fmt.Println("Not sufficient arguments")
		os.Exit(1)
	}

	// Check for subcommand entered
	switch os.Args[1] {

	case "add":
		addCommand.Parse(os.Args[2:])

	case "remove":
		removeCommand.Parse(os.Args[2:])

	default:
		fmt.Println("No Command parsed")
	}

	// Here we are checking if the addCommand is parsed
	if addCommand.Parsed() {
		fmt.Println("add subcommand entered")
		if *userName != "" {
			fmt.Println("Username:", *userName)
		} else {
			fmt.Println("No username entered")
		}
	}

	// Here we are checking if removeCommand is parsed
	if removeCommand.Parsed() {
		fmt.Println("remove subcommand entered")
		if *removeDirectory == true {
			fmt.Println("Home directory of user is removed")
		} else {
			fmt.Println("Home directory of user is not removed")
		}
	}

}
