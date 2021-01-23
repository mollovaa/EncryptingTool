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

var ciphers = map[int]string{
	1: "Ceaser",
	2: "Route",
	3: "Transposition",
	4: "Monoalphabetic",
	5: "One-time pad"}

type OperationEnum int

const (
	Encrypt OperationEnum = 1
	Decrypt OperationEnum = 2
)

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

var statistics = map[string]int{
	"Ceaser":         0,
	"Route":          0,
	"Transposition":  0,
	"Monoalphabetic": 0,
	"One-time pad":   0}

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
