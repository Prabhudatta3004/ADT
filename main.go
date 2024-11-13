package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/prabhudatta3004/ADT/config"
	"github.com/prabhudatta3004/ADT/internal/storage"
)

func main() {
	// Load configuration
	cfg := config.LoadConfig()

	// Initialize the key-value store with the configured data file path
	store := storage.NewKeyValueStore(cfg.DataPath)

	// Create a scanner for user input
	scanner := bufio.NewScanner(os.Stdin)

	// Interactive loop for user commands
	for {
		fmt.Print("Enter command (set/get/delete/exit): ")
		if !scanner.Scan() {
			fmt.Println("\nError reading input or input stream closed. Exiting...")
			break
		}

		// Trim and split the input
		input := strings.TrimSpace(scanner.Text())
		if input == "" {
			continue
		}
		parts := strings.Fields(input)
		command := strings.ToLower(parts[0])

		// Command handling
		switch command {
		case "set":
			if len(parts) < 3 {
				fmt.Println("Usage: set <key> <value>")
				continue
			}
			key := parts[1]
			// Join remaining parts as the value to support spaces in values
			value := strings.Join(parts[2:], " ")
			store.Set(key, value)
			fmt.Printf("Key '%s' set successfully.\n", key)
		case "get":
			if len(parts) < 2 {
				fmt.Println("Usage: get <key>")
				continue
			}
			key := parts[1]
			if value, exists := store.Get(key); exists {
				fmt.Printf("Value for key '%s': %v\n", key, value)
			} else {
				fmt.Printf("Key '%s' not found.\n", key)
			}
		case "delete":
			if len(parts) < 2 {
				fmt.Println("Usage: delete <key>")
				continue
			}
			key := parts[1]
			store.Delete(key)
			fmt.Printf("Key '%s' deleted successfully.\n", key)
		case "exit":
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Unknown command. Available commands: set, get, delete, exit.")
		}
	}
}
