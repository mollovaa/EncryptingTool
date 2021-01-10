package monoalphabetic

import "strings"

const alphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
const alphabetReversed = "ZYXWVUTSRQPONMLKJIHGFEDCBA"

func Encrypt(plaintext string, key string) string {
	cipherPad := makeCipher(key)
	encrypted := ""

	plaintext = strings.ToUpper(plaintext)

	for i := 0; i < len(plaintext); i++ {
		idx := strings.Index(alphabet, string(plaintext[i]))
		encrypted += string(cipherPad[idx])
	}
	return encrypted
}

func Decrypt(ciphertext string, key string) string {
	cipherPad := makeCipher(key)
	decrypted := ""

	for i := 0; i < len(ciphertext); i++ {
		idx := strings.Index(cipherPad, string(ciphertext[i]))
		decrypted += string(alphabet[idx])
	}

	return decrypted
}

func makeCipher(key string) string {

	key = strings.ToUpper(key)

	cipherPad := ""
	for i := 0; i < len(key); i++ {
		if strings.Contains(cipherPad, string(key[i])) == false {
			cipherPad += string(key[i])
		}
	}

	for i := 0; i < len(alphabetReversed); i++ {
		if strings.Contains(cipherPad, string(alphabetReversed[i])) == false {
			cipherPad += string(alphabetReversed[i])
		}
	}

	return cipherPad

}
