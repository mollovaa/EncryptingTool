package caesar

const lowerCaseAlphabet = "abcdefghijklmnopqrstuvwxyz"
const upperCaseAlphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

func Encrypt(plaintext string, key int) string {
	return rotateText(plaintext, key)
}

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
