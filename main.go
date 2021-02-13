/*
	Main package:
	contains the console app logic: use choices, input and output.
*/
package main

import (
	"ceaser"
	"encrypting"
	"monoalphabetic"
	"onetimepad"
	"route"
	"transposition"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

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
// the cipher encryption of input plain text is called.
// The result is the cipher text.
func callCipherEncryption(cipher encrypting.Cipher, scanner *bufio.Scanner) string {
	fmt.Println("Enter plain text to encypt:")
	scanner.Scan()
	plaintext := scanner.Text()

	if cipher.Name() == "Monoalphabetic" || cipher.Name() == "OneTimePad" {
		key := defineTextCipherKey(scanner)
		cipher.IncreaseTimesUsed()
		return cipher.Encrypt(plaintext, key)
	}

	key := defineNumberCipherKey(scanner)
	cipher.IncreaseTimesUsed()
	return cipher.Encrypt(plaintext, strconv.Itoa(key))
}

// Depending on the cipher used,
// the cipher decryption of input cipher text is called.
// The result is the plain text.
func callCipherDecryption(cipher encrypting.Cipher, scanner *bufio.Scanner) string {
	fmt.Println("Enter cipher text to decrypt:")
	scanner.Scan()
	ciphertext := scanner.Text()

	if cipher.Name() == "Monoalphabetic" || cipher.Name() == "OneTimePad" {
		key := defineTextCipherKey(scanner)
		cipher.IncreaseTimesUsed()
		return cipher.Decrypt(ciphertext, key)
	}
	key := defineNumberCipherKey(scanner)
	cipher.IncreaseTimesUsed()
	return cipher.Decrypt(ciphertext, strconv.Itoa(key))
}

func main() {
	ceaserCipher := &ceaser.Ceaser{0}
	monoalphabeticCipher := &monoalphabetic.MonoAlphabetic{0}
	oneTimePadCipher := &onetimepad.OneTimePad{0}
	routeCipher := &route.Route{0}
	transpositionCipher := &transposition.Transposition{0}

	var ciphersMap = map[int]encrypting.Cipher{
		1: ceaserCipher,
		2: monoalphabeticCipher,
		3: oneTimePadCipher,
		4: routeCipher,
		5: transpositionCipher}

	scanner := bufio.NewScanner(os.Stdin)
	var cipher encrypting.Cipher
	var err error

	for {
		encrypting.PrintAvailableCiphers(ciphersMap)

		fmt.Println("Choose the cipher you want to use by its corresponding number.")

		for {
			scanner.Scan()

			cipher, err = encrypting.DefineCipherChoice(ciphersMap, scanner.Text())
			if err == nil {
				fmt.Println("\nCipher chosen:", cipher.Name())
				break
			}
			fmt.Println("Invalid choice. Please try again.")
		}

		fmt.Println("Press 1 for encryption and 2 for decryption.")

		var operation encrypting.OperationEnum
		for {
			scanner.Scan()

			operation, err = encrypting.DefineOperationChoice(scanner.Text())

			if err == nil {
				fmt.Println("\nOperation chosen:", operation)
				break
			}
			fmt.Println("Invalid choice. Please try again.")
		}

		var result string
		switch operation {
		case encrypting.Encrypt:
			result = callCipherEncryption(cipher, scanner)
			break
		case encrypting.Decrypt:
			result = callCipherDecryption(cipher, scanner)
			break
		}
		fmt.Println("Result:", result)

		fmt.Println("Do you want to continue? Choose Y for Yes and anything else for No.")
		scanner.Scan()
		continueChoice := scanner.Text()
		if continueChoice != "Y" && continueChoice != "y" {
			fmt.Println("Statistics:", ciphersMap)
			os.Exit(3)
		}
	}
}
