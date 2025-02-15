package utils

import (
	"fmt"
	"os/exec"
)

func EncryptFile(filename, recipient string) error {
	cmd := exec.Command("gpg", "--encrypt", "--recipient", recipient, "--output", filename+".gpg", filename)
	
	err := cmd.Run()
	if err != nil {
		return err
	}

	fmt.Println("File encrypted:", filename+".gpg")
	return nil
}

func DecryptFile(encryptedFile, outputFile string) error {
	cmd := exec.Command("gpg", "--decrypt", "--output", outputFile, encryptedFile)
	
	err := cmd.Run()
	if err != nil {
		return err
	}

	fmt.Println("File decrypted to:", outputFile)
	return nil
}
