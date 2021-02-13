/*
	Onetimepad package contains implementations of
	encryption and decryption processes with onetimepad cipher.
	To encrypt or decrypt a key is needed.
	The key is a text - only letters, no special symbols or digits.
*/
package onetimepad

import "strings"
import "fmt"

type OneTimePad struct {
	TimesUsed int
}

const alphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
const alphabetSize = 26
const symbol = "X"

// Given are the plain text and the key.
// The key should be with the same size as the plain text,
// if that's not the case, then "X" are appended to the key.
// Each letter from the plain text is summed up with the corresponding one from the key
// (their indices in the alphabet).
// The resulted sum is the alphabet position of the ciphered letter.
func (o *OneTimePad) Encrypt(plaintext string, key string) string {
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

// Given are the cipher text and the key.
// The key should be with the same size as the plain text,
// if that's not the case, then "X" are appended to the key.
// From each cipher text letter, the corresponding key letter is subtracted (their indices in the alphabet).
// The result is the index of the alphabter of the relevant plain text letter.
func (o *OneTimePad) Decrypt(ciphertext string, key string) string {
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

func (o *OneTimePad) String() string {
	return fmt.Sprintf("Name: OneTimePad, used %d times.", o.TimesUsed)
}

func (o *OneTimePad) Name() string {
	return "OneTimePad"
}

func (o *OneTimePad) GetTimesUsed() int {
	return o.TimesUsed
}

func (o *OneTimePad) IncreaseTimesUsed() {
	o.TimesUsed = o.TimesUsed + 1
}
