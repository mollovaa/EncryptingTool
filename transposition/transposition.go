package transposition

import "strings"
import "math"

func Encrypt(plaintext string, key int) string {
	key = int(math.Abs(float64(key)))

	ciphertext := make([]string, len(plaintext))

	for to, from := range createTransition(len(plaintext), key) {
		ciphertext[to] = string(plaintext[from])
	}

	return strings.Join(ciphertext, "")
}

func Decrypt(ciphertext string, key int) string {
	key = int(math.Abs(float64(key)))

	plaintext := make([]string, len(ciphertext))

	// from is the index of the result, to is the index of the ciphertext
	for from, to := range createTransition(len(ciphertext), key) {
		plaintext[to] = string(ciphertext[from])
	}

	return strings.Join(plaintext, "")
}

func createTransition(items int, key int) []int {
	var matrix [][]int
	// The row will be increased/decreased by the coordinate
	// depending on wheter the direction of the diagonal is downwards or upwards
	row := 0
	coordinate := 1

	// Iterating the plain text indices
	for i := 0; i < items; i++ {
		// Add new row to the matrix
		matrix = append(matrix, []int{})

		// Add current index to the defined row
		matrix[row] = append(matrix[row], i)

		// Check if the row exceedes the given depth or if it will be <0 at the next iteration.
		// Make the zig-zag move
		if row+coordinate < 0 || key <= row+coordinate {
			coordinate = -coordinate
		}

		// Increment the row by the coord (which can be negative.)
		row += coordinate
	}

	result := []int{}
	for _, slice := range matrix {
		for _, elem := range slice {
			result = append(result, elem)
		}
	}
	return result
}
