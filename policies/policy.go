package policies

import (
	"regexp"
	"strings"
)

// PasswordRules defines the validation logic for a password
type PasswordRules struct {
	MinLength           int
	MaxLength           int
	UppercaseRequired   bool
	LowercaseRequired   bool
	DigitRequired       bool
	SpecialCharRequired bool
	ExcludeDictionary   bool
}

// ValidatePassword validates the password against given rules
func ValidatePassword(password string, rules PasswordRules) bool {
	// Check length
	if len(password) < rules.MinLength || len(password) > rules.MaxLength {
		return false
	}

	// Check for uppercase
	if rules.UppercaseRequired && !hasUppercase(password) {
		return false
	}

	// Check for lowercase
	if rules.LowercaseRequired && !hasLowercase(password) {
		return false
	}

	// Check for digit
	if rules.DigitRequired && !hasDigit(password) {
		return false
	}

	// Check for special character
	if rules.SpecialCharRequired && !hasSpecialChar(password) {
		return false
	}

	// Check for dictionary word (simplified check for now)
	if rules.ExcludeDictionary && isDictionaryWord(password) {
		return false
	}

	return true
}

func hasUppercase(password string) bool {
	return regexp.MustCompile(`[A-Z]`).MatchString(password)
}

func hasLowercase(password string) bool {
	return regexp.MustCompile(`[a-z]`).MatchString(password)
}

func hasDigit(password string) bool {
	return regexp.MustCompile(`[0-9]`).MatchString(password)
}

func hasSpecialChar(password string) bool {
	return regexp.MustCompile(`[!@#\$%\^&\*\(\)_\+\-=\[\]\{\};:\'",<>\./?\\|]`).MatchString(password)
}

func isDictionaryWord(password string) bool {
	// Implement actual dictionary check, for now a simple check
	if strings.Contains(password, "password") {
		return true
	}
	return false
}
