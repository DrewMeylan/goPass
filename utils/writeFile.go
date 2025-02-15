package utils

import (
	"fmt"
	"os"
)

func WriteFile(filename, password string) error { //expand eventually to include DB name as input
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(password)
	if err != nil {
		return err
	}

	fmt.Println("Password written to file: ", filename)
	return nil
}
