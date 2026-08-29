package main

import (
	"fmt"
	"os"
)
 const version = "0.1.0"

func main() {
	if len(os.Args) < 2 {
		printHelp()
		return 
	}


 command := os.Args[1]

switch command {

case "help":
	printHelp()

case "version":
	printVersion()

default:
	fmt.Printf("Unknown command: %s\n", command)
	printHelp()
}
}

func printHelp() {
	fmt.Println(`YOSI - source-based package manager

Usage:
  yosi <command>

Commands:
  help       Show this help message
  version    Show YOSI version
`)
}

func printVersion() {
	fmt.Printf("YOSI %s\n", version)
}