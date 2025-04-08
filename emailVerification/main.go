package main

import (
	"fmt"
	"net/mail"
	"regexp"
	"strings"
)

func main() {
	fmt.Print("Enter email address to verify: ")
	var email string
	fmt.Scanln(&email)
	email = strings.TrimSpace(email)

	if isEmailValid(email) {
		fmt.Printf("The email address '%s' is valid\n", email)
	} else {
		fmt.Printf("The email address '%s' is NOT valid\n", email)
	}
}

func isEmailValid(email string) bool {
	// Method 1: Using net/mail parser
	_, err := mail.ParseAddress(email)
	if err != nil {
		return false
	}

	// Method 2: Using regex for additional validation
	// Check email format: username@domain.tld
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	if !emailRegex.MatchString(email) {
		return false
	}

	// Additional checks
	if len(email) > 254 { // RFC 5321
		return false
	}

	parts := strings.Split(email, "@")
	if len(parts) != 2 {
		return false
	}

	username := parts[0]
	domain := parts[1]

	// Check username length
	if len(username) > 64 { // RFC 5321
		return false
	}

	// Check for consecutive dots
	if strings.Contains(email, "..") {
		return false
	}

	// Check domain
	if strings.HasPrefix(domain, ".") || strings.HasSuffix(domain, ".") {
		return false
	}

	return true
}
