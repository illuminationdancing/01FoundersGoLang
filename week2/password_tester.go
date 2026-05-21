package main

import (
	"fmt"
	"unicode"
	"strings"
)

func main() {
	var password string

	fmt.Print("Enter a password: ")
	fmt.Scan(&password)

	var (
		hasUpper   bool
		hasLower   bool
		hasNumber  bool
		hasSpecial bool
	)

	blacklist := []string{
		"password", "123456", "qwerty",
		"letmein", "welcome", "monkey",
		"dragon", "master", "sunshine",
	}

	for _, char := range password {
		switch {
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsNumber(char):
			hasNumber = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			hasSpecial = true
		}
	}

	isBlacklisted := false
	for _, banned := range blacklist {
		if strings.Contains(strings.ToLower(password),banned) {
			isBlacklisted = true
		}
	}

	isLongEnough := len(password) >= 8

	score := 0
	if isLongEnough {
		score++
	}
	if hasUpper {
		score++
	}
	if hasLower {
		score++
	}
	if hasNumber {
		score++
	}
	if hasSpecial {
		score++
	}

	fmt.Println("\nI appreciate you trying to create ANOTHER password you'll need to remember.")
	fmt.Println("The human condition is NOT designed for password generation. But tech is.")
	fmt.Println("Here's my feedback on your password.\n")
	fmt.Println("--- Password Strength Report ---")

	if isBlacklisted {
		fmt.Println("❌ Password is too common and easily guessed")
		fmt.Println("Strength: DANGEROUSLY WEAK. YOU WILL BE HACKED 💀")
		return
	}

	if isLongEnough {
		fmt.Println("✅ Length good")
	} else {
		fmt.Println("❌ Too short. Can you count to 8? There's an 8 char minimum for your own security, human.")
	}
	if hasUpper {
		fmt.Println("✅ Has uppercase letters")
	} else {
		fmt.Println("❌ No uppercase letters. Put an UPPERCASE LETTER")
	}
	if hasLower {
		fmt.Println("✅ Has lowercase letters")
	} else {
		fmt.Println("❌ No lowercase letters. Put a lowercase letter.")
	}
	if hasNumber {
		fmt.Println("✅ Has numbers")
	} else {
		fmt.Println("❌ No numbers. Put a number. NOT your grandma's birthday!")
	}
	if hasSpecial {
		fmt.Println("✅ Has special characters")
	} else {
		fmt.Println("❌ No special characters. Special characters are $$$")
	}

	switch {
	case score == 5:
		fmt.Println("Strength: STRONG. UNLIKELY TO BE HACKED 🛡️")
	case score >= 3:
		fmt.Println("Strength: MEDIUM ⚠️")
	default:
		fmt.Println("Strength: WEAK 💀")
	}
}
