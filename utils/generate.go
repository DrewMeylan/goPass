package utils

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"regexp"
	"strings"
)

func generatePassword(length int, requirements map[string]bool) (string, error) {
	// Define character sets
	lowercase := "abcdefghijklmnopqrstuvwxyz"
	uppercase := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	numbers := "0123456789"
	special := "!@#$%^&*()-_=+[]{}|;:,.<>?/"

	var charSet string
	if requirements["lowercase"] {
		charSet += lowercase
	}
	if requirements["uppercase"] {
		charSet += uppercase
	}
	if requirements["numbers"] {
		charSet += numbers
	}
	if requirements["special"] {
		charSet += special
	}

	// Generate the password
	var password strings.Builder
	for i := 0; i < length; i++ {
		idx, err := rand.Int(rand.Reader, big.NewInt(int64(len(charSet))))
		if err != nil {
			return "", err
		}
		password.WriteByte(charSet[idx.Int64()])
	}

	return password.String(), nil
}

// Example of password policy enforcement
func validatePassword(password string, policy map[string]bool) bool {
	if policy["uppercase"] && !regexp.MustCompile(`[A-Z]`).MatchString(password) {
		return false
	}
	if policy["lowercase"] && !regexp.MustCompile(`[a-z]`).MatchString(password) {
		return false
	}
	if policy["numbers"] && !regexp.MustCompile(`[0-9]`).MatchString(password) {
		return false
	}
	if policy["special"] && !regexp.MustCompile(`[!@#$%^&*()_+\-=\[\]{};':"\\|,.<>\/?]`).MatchString(password) {
		return false
	}
	return true
}

func main() {
	// Define password policy
	requirements := map[string]bool{
		"lowercase": true,
		"uppercase": true,
		"numbers":   true,
		"special":   true,
	}

	// Generate a password based on the policy
	password, err := generatePassword(12, requirements)
	if err != nil {
		fmt.Println("Error generating password:", err)
		return
	}

	// Validate password
	if validatePassword(password, requirements) {
		fmt.Println("Generated password is valid:", password)
	} else {
		fmt.Println("Generated password does not meet the policy.")
	}
}
