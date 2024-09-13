package main

import (
	"fmt"
	"os"

	utils "github.com/null-none/password_manager/utils"
)

const filename = "passwords.json"

func main() {
	fmt.Println("Welcome to the Password Manager!")
	var entries []utils.PasswordEntry
	// Load existing passwords from file
	entries, err := utils.LoadPasswordsFromFile(filename)
	if err != nil {
		fmt.Println("Error loading passwords:", err)
	}
	for {
		fmt.Println("\nChoose an option:")
		fmt.Println("1. Generate a new password")
		fmt.Println("2. View saved passwords")
		fmt.Println("3. Exit")
		var choice int
		fmt.Scanln(&choice)
		switch choice {
		case 1:
			entry := getPasswordEntryFromUser()
			entries = append(entries, entry)
			if err := utils.SavePasswordsToFile(entries, filename); err != nil {
				fmt.Println("Error saving password:", err)
			}
		case 2:
			displayPasswords(entries)
		case 3:
			fmt.Println("Thank you for using Password Manager!")
			os.Exit(0)
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}
func getPasswordEntryFromUser() utils.PasswordEntry {
	var service, username string
	var length int
	var useUppercase, useDigits, useSpecialChars bool
	fmt.Print("Enter the service name: ")
	fmt.Scanln(&service)
	fmt.Print("Enter the username: ")
	fmt.Scanln(&username)
	fmt.Print("Enter the desired password length: ")
	fmt.Scanln(&length)
	fmt.Print("Use uppercase letters? (Y/N): ")
	fmt.Scanln(&useUppercase)
	fmt.Print("Use digits? (Y/N): ")
	fmt.Scanln(&useDigits)
	fmt.Print("Use special characters? (Y/N): ")
	fmt.Scanln(&useSpecialChars)
	pass := utils.GeneratePassword(length, useUppercase, useDigits, useSpecialChars)
	return utils.PasswordEntry{
		Service:  service,
		Username: username,
		Password: pass,
	}
}
func displayPasswords(entries []utils.PasswordEntry) {
	if len(entries) == 0 {
		fmt.Println("No passwords saved yet.")
		return
	}
	fmt.Println("\nSaved Passwords:")
	for _, entry := range entries {
		fmt.Printf("Service: %s | Username: %s | Password: %s\n", entry.Service, entry.Username, entry.Password)
	}
}
