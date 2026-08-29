package main

import (
	"fmt"
	"os"
	"yosi/internal/repo"
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

	case "search":
		if len(os.Args) < 3 {
			fmt.Println("Usage: yosi search <package-name>")
			return
		}

		packageName := os.Args[2]

		index, err := repo.FetchIndex()
		if err != nil {
			fmt.Printf("Error fetching repository index: %v\n", err)
			return
		}
		for _, pkg := range index.Packages {
			if pkg.Name == packageName {
				fmt.Printf("%s %s\n", pkg.Name, pkg.Version)
				return
			}
		}
		fmt.Printf("Package %s not found\n", packageName)

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
  search     Search for a package`)
}

func printVersion() {
	fmt.Printf("YOSI %s\n", version)
}
