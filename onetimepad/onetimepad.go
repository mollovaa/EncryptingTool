package onetimepad

import "strings"

const alphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
const alphabetSize = 26
const symbol = "X"

func Encrypt(plaintext string, key string) string {
	plaintext = strings.ToUpper(plaintext)
	key = strings.ToUpper(key)

	if len(key) < len(plaintext) {
		for i := 0; i < len(plaintext)-len(key); i++ {
			key += symbol
		}
	}

	ciphertext := ""

	for i := 0; i < len(plaintext); i++ {
		idx1 := strings.Index(alphabet, string(plaintext[i]))
		idx2 := strings.Index(alphabet, string(key[i]))
		elementIdx := (idx1 + idx2) % alphabetSize

		ciphertext += string(alphabet[elementIdx])
	}

	return ciphertext
}

func Decrypt(ciphertext string, key string) string {
	ciphertext = strings.ToUpper(ciphertext)
	key = strings.ToUpper(key)

	if len(key) < len(ciphertext) {
		for i := 0; i < len(ciphertext)-len(key); i++ {
			key += symbol
		}
	}

	plaintext := ""

	for i := 0; i < len(ciphertext); i++ {

		idx1 := strings.Index(alphabet, string(ciphertext[i]))
		idx2 := strings.Index(alphabet, string(key[i]))

		var elementIdx int
		if idx1-idx2 < 0 {
			elementIdx = (idx1 - idx2 + alphabetSize) % alphabetSize
		} else {
			elementIdx = (idx1 - idx2) % alphabetSize
		}

		plaintext += string(alphabet[elementIdx])
	}
	return plaintext
}
