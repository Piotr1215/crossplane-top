package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	// Define flags
	var name string
	flag.StringVar(&name, "name", "World", "a name to say hello to")

	// Parse the flags
	flag.Parse()

	// Main logic of the CLI
	if name == "" {
		fmt.Println("Please provide a name")
		os.Exit(1)
	}

	fmt.Printf("Hello, %s!\n", name)
}
