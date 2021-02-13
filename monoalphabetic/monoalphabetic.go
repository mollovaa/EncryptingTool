/*
	Monoalphabetic package contains implementations of
	encryption and decryption processes with monoalphabetic cipher.
	To encrypt or decrypt a key is needed.
	The key is a text - only letters, no special symbols or digits.
	The operatons are done with capital letters only.
*/
package monoalphabetic

import "strings"
import "fmt"

type MonoAlphabetic struct {
	TimesUsed int
}

const alphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
const alphabetReversed = "ZYXWVUTSRQPONMLKJIHGFEDCBA"

// Given are the plain text and the key.
// Duplicated letters are removed from the key
// and the letters from the alphabet which are not present,
// are appended to the key in reversed order (cipher pad).
// The cipher text is structured using the cipher pad,
// the letters corresponding to the plaintext's letter position, are the relevant ones.
func (m *MonoAlphabetic) Encrypt(plaintext string, key string) string {
	cipherPad := makeCipher(key)
	encrypted := ""

	plaintext = strings.ToUpper(plaintext)

	for i := 0; i < len(plaintext); i++ {
		idx := strings.Index(alphabet, string(plaintext[i]))
		encrypted += string(cipherPad[idx])
	}
	return encrypted
}

// Given are the cipher text and the key.
// The cipher pad is created based on the key and the plain text,
// letters are taken from the alphabet based on the position of the cipher text letters in the cipher pad.
func (m *MonoAlphabetic) Decrypt(ciphertext string, key string) string {
	cipherPad := makeCipher(key)
	decrypted := ""

	for i := 0; i < len(ciphertext); i++ {
		idx := strings.Index(cipherPad, string(ciphertext[i]))
		decrypted += string(alphabet[idx])
	}

	return decrypted
}

func (c *MonoAlphabetic) String() string {
	return fmt.Sprintf("Name: Monoalphabetic, used %d times.", c.TimesUsed)
}

func (m *MonoAlphabetic) Name() string {
	return "Monoalphabetic"
}

func (m *MonoAlphabetic) GetTimesUsed() int {
	return m.TimesUsed
}

func (m *MonoAlphabetic) IncreaseTimesUsed() {
	m.TimesUsed = m.TimesUsed + 1
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
