package main

import (
	"fmt"
	"os"
	"strings"

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

		query := strings.ToLower(strings.TrimSpace(packageName))
		matches := 0

		for _, pkg := range index.Packages {
			if strings.Contains(strings.ToLower(pkg.Name), query) {
				fmt.Printf("%-20s %s\n", pkg.Name, pkg.Version)
				matches++
			}
		}

		if matches == 0 {
			fmt.Printf("Package '%s' not found in the repository.\n", packageName)
		}

	case "info":
		if len(os.Args) < 3 {
			fmt.Println("Usage: yosi info <package-name>")
			return
		}

		packageName := os.Args[2]

		index, err := repo.FetchIndex()
		if err != nil {
			fmt.Printf("Error fetching repository index: %v\n", err)
			return
		}

		for _, pkg := range index.Packages {
			if strings.EqualFold(pkg.Name, packageName) {
				manifest, err := repo.FetchManifest(pkg.Name, pkg.Version)
				if err != nil {
					fmt.Printf("Error fetching package manifest: %v\n", err)
					return
				}

				fmt.Printf("Name: %s\n", manifest.Package.Name)
				fmt.Printf("Version: %s\n", manifest.Package.Version)
				fmt.Printf("Description: %s\n", manifest.Package.Description)
				fmt.Printf("License: %s\n", manifest.Package.License)
				fmt.Printf("Homepage: %s\n", manifest.Package.Homepage)
				fmt.Printf("Maintainer: %s\n", manifest.Package.Maintainer)
				fmt.Printf(
					"Architecture: %s\n",
					strings.Join(manifest.Metadata.Architectures, ", "),
				)
				return
			}
		}

		fmt.Printf("Package '%s' not found in the repository.\n", packageName)

	default:
		fmt.Printf("Unknown command: %s\n\n", command)
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
  search     Search for a package
  info       Show package information`)
}

func printVersion() {
	fmt.Printf("YOSI %s\n", version)
}
