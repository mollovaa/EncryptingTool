/*
	Ceaser package contains implementations of
	encryption and decryption processes with ceaser cipher.
	To encrypt or decrypt a key is needed.
	The key is a number - might be positive and negative.
	Only letters are encrypted and decrypted.
*/
package ceaser

import "strconv"
import "fmt"

type Ceaser struct {
	TimesUsed int
}

const allAsciiCharacters = "!\"#$%&'()*+,-./0123456789:;<=>?@ABCDEFGHIJKLMNOPQRSTUVWXYZ[\\]^_`abcdefghijklmnopqrstuvwxyz"
const size = len(allAsciiCharacters)

// Given are the plain text and the key.
// Each letter of the plain text is shifted "key" position to left or right
// (depending on whether the number is positive or negative) in the alphabet.
// The result is the cipher text.
func (c *Ceaser) Encrypt(plaintext string, key string) string {
	i, _ := strconv.Atoi(key)
	return rotateText(plaintext, i)
}

// Given are the cipher text and the key.
// Each letter of the plain text is shifted to the opposite direction of the one used for encryption.
// The result is the plain text.
func (c *Ceaser) Decrypt(ciphertext string, key string) string {
	i, _ := strconv.Atoi(key)
	return rotateText(ciphertext, -i)
}

func (c *Ceaser) String() string {
	return fmt.Sprintf("Name: Ceaser, used %d times.", c.TimesUsed)
}

func (c *Ceaser) Name() string {
	return "Ceaser"
}

func (c *Ceaser) GetTimesUsed() int {
	return c.TimesUsed
}

func (c *Ceaser) IncreaseTimesUsed() {
	c.TimesUsed = c.TimesUsed + 1
}

func rotateText(inputText string, key int) string {
	// If the key is > 90, its = key % 90.
	key %= size
	rotatedText := []byte(inputText)

	for index, byteValue := range rotatedText {
			shiftedIdx := (int((size + int(byteValue - '!'))) + key) % size
			rotatedText[index] = allAsciiCharacters[shiftedIdx]
	}
	return string(rotatedText)
}
