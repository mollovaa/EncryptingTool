/*
	Main package:
	contains the console app logic: use choices, input and output.
*/
package main

import (
	ceaser "./ceaser"
	monoalphabetic "./monoalphabetic"
	onetimepad "./onetimepad"
	route "./route"
	transposition "./transposition"
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"unicode"
)

// Map ciphers stores all available ciphers
var ciphers = map[int]string{
	1: "Ceaser",
	2: "Route",
	3: "Transposition",
	4: "Monoalphabetic",
	5: "One-time pad"}

// Custom type OperationEnum with 2 options: Encrypt, Decrypt
type OperationEnum int

const (
	Encrypt OperationEnum = 1
	Decrypt OperationEnum = 2
)

// stringify function of the OperationEnum
func (e OperationEnum) String() string {
	switch e {
	case Encrypt:
		return "Encryption"
	case Decrypt:
		return "Decryption"
	default:
		return "InvalidOperation"
	}
}

// Function which prints all available ciphers with their corresponding keys in the map
func printAvailableCiphers() {
	var keys []int
	for key, _ := range ciphers {
		keys = append(keys, key)
	}
	sort.Ints(keys)

	fmt.Println("Available ciphers:")
	for _, key := range keys {
		fmt.Println(key, ":", ciphers[key])
	}
}

// Reads input from the console.
// Validates if the input is a corresponding number to a cipher.
// If yes, the number is returned, else error is thrown.
func defineCipherChoice(scanner *bufio.Scanner) int {
	fmt.Println("Choose the cipher you want to use by its corresponding number.")
	for {
		scanner.Scan()
		cipher, err := strconv.Atoi(scanner.Text())

		if err == nil && cipher >= 1 && cipher <= len(ciphers) {
			return cipher
		}
		fmt.Println("Invalid choice. Please try again.")
	}
}

// Reads input from the console.
// Validates if the input is a number.
// If yes, the number, used later for cipher key, is returned, else error is thrown.
func defineNumberCipherKey(scanner *bufio.Scanner) int {
	fmt.Println("Enter the key used for encryption. It should be a number:")
	for {
		scanner.Scan()
		key, err := strconv.Atoi(scanner.Text())

		if err == nil {
			return key
		}
		fmt.Println("Invalid input. Please try again.")
	}
}

// Reads input from the console.
// Validates if the input is a number corresponding to a operarion.
// If yes, the operationEnum is returned, else error is thrown.
func validateOperationChoice(scanner *bufio.Scanner) OperationEnum {
	fmt.Println("Press 1 for encryption and 2 for decryption.")
	for {
		scanner.Scan()
		option, err := strconv.Atoi(scanner.Text())

		if err == nil && (option == 1 || option == 2) {
			switch option {
			case 1:
				return Encrypt
			case 2:
				return Decrypt
			}
		}
		fmt.Println("Invalid choice. Please try again.")
	}
}

// Stores statistics about which cipher is the most used.
var statistics = map[string]int{
	"Ceaser":         0,
	"Route":          0,
	"Transposition":  0,
	"Monoalphabetic": 0,
	"One-time pad":   0}

// Depending on the cipher used, 
// the cipher encryption of input plain text is called.
// The result is the cipher text.
func callCipherEncryption(cipherKey int, scanner *bufio.Scanner) string {
	fmt.Println("Enter plain text to encypt:")
	scanner.Scan()
	plaintext := scanner.Text()

	switch cipherKey {
	case 1:
		key := defineNumberCipherKey(scanner)
		statistics["Ceaser"] += 1
		return ceaser.Encrypt(plaintext, key)
	case 2:
		key := defineNumberCipherKey(scanner)
		statistics["Route"] += 1
		return route.Encrypt(plaintext, key)
	case 3:
		key := defineNumberCipherKey(scanner)
		statistics["Transposition"] += 1
		return transposition.Encrypt(plaintext, key)
	case 4:
		key := defineTextCipherKey(scanner)
		statistics["Monoalphabetic"] += 1
		return monoalphabetic.Encrypt(plaintext, key)
	case 5:
		key := defineTextCipherKey(scanner)
		statistics["One-time pad"] += 1
		return onetimepad.Encrypt(plaintext, key)
	default:
		return "InvalidCipherKey!"
	}
}

// Reads input from the console.
// Validates if the input is text only.
// If yes, the text, used later for cipher key, is returned, else error is thrown.
func defineTextCipherKey(scanner *bufio.Scanner) string {
	fmt.Println("Enter the key used for encryption. It should be letters only.")
	for {
		scanner.Scan()
		text := scanner.Text()

		isValid := true
		for _, letter := range text {
			if !unicode.IsLetter(letter) {
				isValid = false
				break
			}
		}

		if isValid == true {
			return text
		}
		fmt.Println("Invalid input. Please try again.")
	}
}

// Depending on the cipher used, 
// the cipher decryption of input cipher text is called.
// The result is the plain text.
func callCipherDecryption(cipherKey int, scanner *bufio.Scanner) string {
	fmt.Println("Enter cipher text to decrypt:")
	scanner.Scan()
	ciphertext := scanner.Text()

	switch cipherKey {
	case 1:
		key := defineNumberCipherKey(scanner)
		statistics["Ceaser"] += 1
		return ceaser.Decrypt(ciphertext, key)
	case 2:
		key := defineNumberCipherKey(scanner)
		statistics["Route"] += 1
		return route.Decrypt(ciphertext, key)
	case 3:
		key := defineNumberCipherKey(scanner)
		statistics["Transposition"] += 1
		return transposition.Decrypt(ciphertext, key)
	case 4:
		key := defineTextCipherKey(scanner)
		statistics["Monoalphabetic"] += 1
		return monoalphabetic.Decrypt(ciphertext, key)
	case 5:
		key := defineTextCipherKey(scanner)
		statistics["One-time pad"] += 1
		return onetimepad.Decrypt(ciphertext, key)
	default:
		return "InvalidCipherKey!"
	}
}

func main() {

	for {
		printAvailableCiphers()

		scanner := bufio.NewScanner(os.Stdin)

		cipher := defineCipherChoice(scanner)

		fmt.Println("\nCipher chosen:", ciphers[cipher])

		operation := validateOperationChoice(scanner)

		fmt.Println("\nOperation chosen:", operation)

		var result string
		switch operation {
		case Encrypt:
			result = callCipherEncryption(cipher, scanner)
			break
		case Decrypt:
			result = callCipherDecryption(cipher, scanner)
			break
		}
		fmt.Println("Result:", result)

		fmt.Println("Do you want to continue? Choose Y for Yes and anything else for No.")
		scanner.Scan()
		continueChoice := scanner.Text()
		if continueChoice != "Y" && continueChoice != "y" {
			fmt.Println("Statistics:", statistics)
			os.Exit(3)
		}
	}
}
