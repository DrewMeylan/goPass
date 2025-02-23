package utils

import (
	"fmt"
	"os"
)

// This work but I have to find a way to get the recipient information from the exec. context
// i.e., from the recipient associated with the database itself. 
func WriteFile(filename, content string, recipient string) error {
	// Write content to file
	err := os.WriteFile(filename, []byte(content), 0600)
	if err != nil {
		return fmt.Errorf("error writing to file: %v", err)
	}

	// Encrypt using GPG
	gpgFile := filename + ".gpg"
	cmd := exec.Command("gpg", "--encrypt", "--recipient", recipient, "--output", gpgFile, filename)
	err = cmd.Run()
	if err != nil {
		return fmt.Errorf("error encrypting file: %v", err)
	}

	// Delete original file
	err = os.Remove(filename)
	if err != nil {
		return fmt.Errorf("error deleting original file: %v", err)
	}

	fmt.Println("File encrypted successfully:", gpgFile)
	return nil
}
