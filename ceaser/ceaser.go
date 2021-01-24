/*
	Ceaser package contains implementations of 
	encryption and decryption processes with ceaser cipher.
	To encrypt or decrypt a key is needed. 
	The key is a number - might be positive and negative.
	Only letters are encrypted and decrypted.
*/
package ceaser

const lowerCaseAlphabet = "abcdefghijklmnopqrstuvwxyz"
const upperCaseAlphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

// Given are the plain text and the key.
// Each letter of the plain text is shifted "key" position to left or right 
// (depending on whether the number is positive or negative) in the alphabet. 
// The result is the cipher text.
func Encrypt(plaintext string, key int) string {
	return rotateText(plaintext, key)
}

// Given are the cipher text and the key.
// Each letter of the plain text is shifted to the opposite direction of the one used for encryption. 
// The result is the plain text.
func Decrypt(ciphertext string, key int) string {
	return rotateText(ciphertext, -key)
}

func rotateText(inputText string, key int) string {
	// If the key is > 26, its = key % 26.
	key %= 26
	rotatedText := []byte(inputText)

	for index, byteValue := range rotatedText {
		if byteValue >= 'a' && byteValue <= 'z' {
			shiftedIdx := (int((26 + (byteValue - 'a'))) + key) % 26
			rotatedText[index] = lowerCaseAlphabet[shiftedIdx]
		} else if byteValue >= 'A' && byteValue <= 'Z' {
			shiftedIdx := (int((26 + (byteValue - 'A'))) + key) % 26
			rotatedText[index] = upperCaseAlphabet[shiftedIdx]
		}
	}
	return string(rotatedText)
}
