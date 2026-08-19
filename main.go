package main

import (
	"fmt"
	"os"
)

// backup_automator - Automated backup scheduling
func backup_automator(path string) {
	fmt.Println("========================================")
	fmt.Println("  Backup-Automator")
	fmt.Println("  Automated backup scheduling")
	fmt.Println("========================================")
	fmt.Println()
	fmt.Println("Target:", path)
	fmt.Println("Processing...")
	fmt.Println("Done!")
}

func main() {
	path := "."
	if len(os.Args) > 1 {
		path = os.Args[1]
	}
	backup_automator(path)
}
