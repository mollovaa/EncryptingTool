/*
	Transposition package contains implementations of
	encryption and decryption processes with transposition cipher.
	To encrypt or decrypt a key is needed.
	The key is a number - no matter positive or negative, because its absolute value is used.
*/
package transposition

import "strings"
import "math"
import "strconv"
import "fmt"

type Transposition struct {
	TimesUsed int
}

// Given are the plain text and the key.
// The plain text is written in diagonals in a zig-zag manner to the selected depth (the key).
// To obtain the cipher text, the diagonals are read row-by-row.
func (t *Transposition) Encrypt(plaintext string, key string) string {
	i, _ := strconv.Atoi(key)
	i = int(math.Abs(float64(i)))

	ciphertext := make([]string, len(plaintext))

	for to, from := range createTransition(len(plaintext), i) {
		ciphertext[to] = string(plaintext[from])
	}

	return strings.Join(ciphertext, "")
}

// Given are the cipher text and the key.
// The cipher text is written in a zig-zag manner to the selected depth (the key).
// To obtain the plain text, the diagonals are read as diagonals (column-by-column).
func (t *Transposition) Decrypt(ciphertext string, key string) string {
	i, _ := strconv.Atoi(key)

	i = int(math.Abs(float64(i)))

	plaintext := make([]string, len(ciphertext))

	// from is the index of the result, to is the index of the ciphertext
	for from, to := range createTransition(len(ciphertext), i) {
		plaintext[to] = string(ciphertext[from])
	}

	return strings.Join(plaintext, "")
}

func (t *Transposition) String() string {
	return fmt.Sprintf("Name: Transposition, used %d times.", t.TimesUsed)
}

func (t *Transposition) Name() string {
	return "Transposition"
}

func (t *Transposition) GetTimesUsed() int {
	return t.TimesUsed
}

func (t *Transposition) IncreaseTimesUsed() {
	t.TimesUsed = t.TimesUsed + 1
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
