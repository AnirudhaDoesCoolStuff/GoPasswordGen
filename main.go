package main

import (
	"crypto/rand"
	"flag"
	"fmt"
	"math/big"
	"strings"
)

func main() {
	length := flag.Int("length", 8, "Length of the password")
	useLower := flag.Bool("lower", true, "Include lowercase letters")
	useUpper := flag.Bool("upper", true, "Include uppercase letters")
	useDigits := flag.Bool("digits", true, "Include digits")
	useSpecial := flag.Bool("special", false, "Include special characters") 
	
	flag.Parse()

	password, err := generatePassword(*length, *useLower, *useUpper, *useDigits,  *useSpecial)
	if err != nil {
		fmt.Println("Error generating the password: ", err)
		return
	}
	fmt.Println("Password:", password)
}

func generatePassword(length int, useLower, useUpper, useDigits, useSpecial bool) (string, error) {
	const lowercase = "abcdefghijklmnopqrstuvwxyz"
	const uppercase = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	const digits = "0123456789"
	const special = "!@#$%^&*()-_=+[]{}|;:,.<>?"

	var charset strings.Builder

	if useLower {
		charset.WriteString(lowercase)
	}
	if useUpper {
		charset.WriteString(uppercase)
	}
	if useDigits {
		charset.WriteString(digits)
	}
	if useSpecial {
		charset.WriteString(special)
	}

	if charset.Len() == 0 {
		charset.WriteString(lowercase)
	}

	finalString := charset.String()

	password := make([]byte, length)

	for i := 0; i < length; i++ {
		randomIndex, err := rand.Int(rand.Reader, big.NewInt(int64(len(finalString))))
		if err != nil {
			return "", err
		}

		password[i] = finalString[randomIndex.Int64()]
	}

	return string(password), nil
}