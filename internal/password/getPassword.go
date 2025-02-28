package password

import (
	"fmt"
	"golang.org/x/term"
	"github.com/pterm/pterm"
  "syscall"
)

func GetPassword() string {
	for {
		// Prompt for first password
		pterm.Info.Println("Enter your password:")
		password1, err := term.ReadPassword(int(syscall.Stdin))
		fmt.Println() // Move to a new line after input
		if err != nil {
			pterm.Error.Println("Error reading password:", err)
			continue
		}

		// Prompt for confirmation
		pterm.Info.Println("Re-enter your password:")
		password2, err := term.ReadPassword(int(syscall.Stdin))
		fmt.Println()
		if err != nil {
			pterm.Error.Println("Error reading password:", err)
			continue
		}

		// Check if passwords match
		if string(password1) == string(password2) {
			pterm.Success.Println("Passwords match!")
			return string(password1)
		} else {
			pterm.Warning.Println("Passwords do not match. Please try again.")
		}
	}
}
