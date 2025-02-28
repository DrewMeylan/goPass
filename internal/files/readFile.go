package files

import (
	"fmt"
	"os"
	"os/exec"
)

// decryptAndReadFile decrypts, prints, and re-encrypts the file.
func ReadFile(filename, recipient string) error {
	decryptedFile := filename + ".dec"

	// Decrypt file
	cmd := exec.Command("gpg", "--decrypt", "--output", decryptedFile, filename)
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("error decrypting file: %v", err)
	}

	// Read and print file content
	data, err := os.ReadFile(decryptedFile)
	if err != nil {
		return fmt.Errorf("error reading decrypted file: %v", err)
	}

	fmt.Println("Decrypted Content:\n", string(data))

	// Re-encrypt the file
	err = WriteFile(decryptedFile, string(data), recipient)
	if err != nil {
		return fmt.Errorf("error re-encrypting file: %v", err)
	}

	// Delete decrypted file
	err = os.Remove(decryptedFile)
	if err != nil {
		return fmt.Errorf("error deleting decrypted file: %v", err) // This throws an error; troubleshooting required
	}

	fmt.Println("File re-encrypted successfully")
	return nil
}

