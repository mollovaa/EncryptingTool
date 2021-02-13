/*
	Route package contains implementations of
	encryption and decryption processes with route cipher.
	To encrypt or decrypt a key is needed.
	The key is a number - might be positive or negative.
*/
package route

import "math"
import "strconv"
import "fmt"

type Route struct {
	TimesUsed int
}

// Given are the plain text and the key.
// The plain text is written in a grid with number of columns = key.
// If some empty cells are left, they are filled with "X", so that a rectangle is formed.
// Next, the letters from the grid are combined following certain route:
// - if the key is a positive number, then the route starts from top left corner and it goes downwards and inwards.
// - if the key is a negative number, then the route starts from the bottom right corner and it goes upwards and inwards.
// The letters in the defined route form the cipher text.
func (r *Route) Encrypt(plaintext string, key string) string {
	i, _ := strconv.Atoi(key)
	columns := int(math.Abs(float64(i)))

	var rows int
	if len(plaintext)%columns == 0 {
		rows = len(plaintext) / columns
	} else {
		rows = len(plaintext)/columns + 1
	}

	grid := createGrid(plaintext, rows, columns)

	if i >= 0 {
		return createCipherTextByPositiveSpiral(grid)
	}
	return createCipherTextByNegativeSpiral(grid)
}

// Given are the cipher text and the key.
// An empty grid is created considering thatthe key is the column number
// and the row number is the lenght of the cipher text divided by the key.
// Then following the route (defined by the sign of the key),
// the letters from the cipher text are placed at the empty grid.
// The plain text is the obtained text from the created grid.
func (r *Route) Decrypt(ciphertext string, key string) string {
	i, _ := strconv.Atoi(key)
	columns := int(math.Abs(float64(i)))
	rows := len(ciphertext) / columns

	grid := createEmptyGrid(rows, columns)

	if i >= 0 {
		fillEmptyGridByPositiveSpiral(grid, ciphertext)
	} else {
		fillEmptyGridByNegativeSpiral(grid, ciphertext)
	}

	plaintext := ""

	for i := 0; i < len(grid); i++ {
		plaintext += grid[i]
	}
	return plaintext

}

func (r *Route) String() string {
	return fmt.Sprintf("Name: Route, used %d times.", r.TimesUsed)
}

func (r *Route) Name() string {
	return "Route"
}

func (r *Route) GetTimesUsed() int {
	return r.TimesUsed
}

func (r *Route) IncreaseTimesUsed() {
	r.TimesUsed = r.TimesUsed + 1
}

func createGrid(plaintext string, rows int, columns int) []string {
	var grid []string
	lastTakenIdx := 0

	for i := 0; i < rows; i++ {
		row := ""

		for j := 0; j < columns; j++ {
			if lastTakenIdx < len(plaintext) {
				row += string(plaintext[lastTakenIdx])
			} else {
				row += "X"
			}
			lastTakenIdx++
		}
		grid = append(grid, row)
	}
	return grid
}

func createCipherTextByPositiveSpiral(grid []string) string {
	gridSize := len(grid) * len(grid[0])
	ciphertext := ""

	var crrMinRow, crrMinColumn int
	var crrMaxRow = len(grid) - 1
	var crrMaxColumn = len(grid[0]) - 1

	for {
		if crrMinRow == crrMaxRow && crrMinColumn == crrMaxColumn {
			break
		}

		// First column
		for i := crrMinRow; i <= crrMaxRow; i++ {
			ciphertext += string(grid[i][crrMinColumn])
		}
		crrMinColumn++
		if len(ciphertext) == gridSize {
			break
		}

		// Last row
		for i := crrMinColumn; i <= crrMaxColumn; i++ {
			ciphertext += string(grid[crrMaxRow][i])
		}
		crrMaxRow--
		if len(ciphertext) == gridSize {
			break
		}

		// Last column upwards:
		for i := crrMaxRow; i >= crrMinRow; i-- {
			ciphertext += string(grid[i][crrMaxColumn])
		}
		crrMaxColumn--
		if len(ciphertext) == gridSize {
			break
		}

		// First row backwards:
		for i := crrMaxColumn; i >= crrMinColumn; i-- {
			ciphertext += string(grid[crrMinRow][i])
		}
		crrMinRow++
		if len(ciphertext) == gridSize {
			break
		}
	}
	return ciphertext
}

func createCipherTextByNegativeSpiral(grid []string) string {
	gridSize := len(grid) * len(grid[0])
	ciphertext := ""

	var crrMinRow, crrMinColumn int
	var crrMaxRow = len(grid) - 1
	var crrMaxColumn = len(grid[0]) - 1

	for {
		if crrMinRow == crrMaxRow && crrMinColumn == crrMaxColumn {
			break
		}

		// Last column upwards:
		for i := crrMaxRow; i >= crrMinRow; i-- {
			ciphertext += string(grid[i][crrMaxColumn])
		}
		crrMaxColumn--
		if len(ciphertext) == gridSize {
			break
		}

		// First row backwards:
		for i := crrMaxColumn; i >= crrMinColumn; i-- {
			ciphertext += string(grid[crrMinRow][i])
		}
		crrMinRow++
		if len(ciphertext) == gridSize {
			break
		}

		// First column
		for i := crrMinRow; i <= crrMaxRow; i++ {
			ciphertext += string(grid[i][crrMinColumn])
		}
		crrMinColumn++
		if len(ciphertext) == gridSize {
			break
		}

		// Last row
		for i := crrMinColumn; i <= crrMaxColumn; i++ {
			ciphertext += string(grid[crrMaxRow][i])
		}
		crrMaxRow--
		if len(ciphertext) == gridSize {
			break
		}
	}
	return ciphertext
}

func createEmptyGrid(rows int, columns int) []string {
	var grid []string

	emptyRow := ""
	for i := 0; i < columns; i++ {
		emptyRow += "_"
	}

	for i := 0; i < rows; i++ {
		grid = append(grid, emptyRow)
	}

	return grid
}

func fillEmptyGridByPositiveSpiral(grid []string, ciphertext string) {

	var lastTakenIdx, crrMinRow, crrMinColumn int
	var crrMaxRow = len(grid) - 1
	var crrMaxColumn = len(grid[0]) - 1

	for {
		if crrMinRow == crrMaxRow && crrMinColumn == crrMaxColumn {
			break
		}

		// Fill First column
		for i := crrMinRow; i <= crrMaxRow; i++ {
			grid[i] = replaceAtIndex(grid[i], ciphertext[lastTakenIdx], crrMinColumn)
			lastTakenIdx++
			// grid[i][crrMinColumn] = ciphertext[lastTakenIdx++]
		}
		crrMinColumn++
		if lastTakenIdx >= len(ciphertext) {
			break
		}

		// Fill Last row
		for i := crrMinColumn; i <= crrMaxColumn; i++ {
			grid[crrMaxRow] =
				replaceAtIndex(grid[crrMaxRow], ciphertext[lastTakenIdx], i)
			lastTakenIdx++
		}
		crrMaxRow--
		if lastTakenIdx >= len(ciphertext) {
			break
		}

		// Fill Last column upwards:
		for i := crrMaxRow; i >= crrMinRow; i-- {
			grid[i] = replaceAtIndex(grid[i], ciphertext[lastTakenIdx], crrMaxColumn)
			lastTakenIdx++
		}
		crrMaxColumn--
		if lastTakenIdx >= len(ciphertext) {
			break
		}

		// Fill First row backwards:
		for i := crrMaxColumn; i >= crrMinColumn; i-- {
			grid[crrMinRow] = replaceAtIndex(grid[crrMinRow], ciphertext[lastTakenIdx], i)
			lastTakenIdx++
		}
		crrMinRow++
		if lastTakenIdx >= len(ciphertext) {
			break
		}
	}

}

func fillEmptyGridByNegativeSpiral(grid []string, ciphertext string) {

	var lastTakenIdx, crrMinRow, crrMinColumn int
	var crrMaxRow = len(grid) - 1
	var crrMaxColumn = len(grid[0]) - 1

	for {
		if crrMinRow == crrMaxRow && crrMinColumn == crrMaxColumn {
			break
		}

		// Fill Last column upwards:
		for i := crrMaxRow; i >= crrMinRow; i-- {
			grid[i] = replaceAtIndex(grid[i], ciphertext[lastTakenIdx], crrMaxColumn)
			lastTakenIdx++
		}
		crrMaxColumn--
		if lastTakenIdx >= len(ciphertext) {
			break
		}

		// Fill First row backwards:
		for i := crrMaxColumn; i >= crrMinColumn; i-- {
			grid[crrMinRow] = replaceAtIndex(grid[crrMinRow], ciphertext[lastTakenIdx], i)
			lastTakenIdx++
		}
		crrMinRow++
		if lastTakenIdx >= len(ciphertext) {
			break
		}

		// Fill First column
		for i := crrMinRow; i <= crrMaxRow; i++ {
			grid[i] = replaceAtIndex(grid[i], ciphertext[lastTakenIdx], crrMinColumn)
			lastTakenIdx++
			// grid[i][crrMinColumn] = ciphertext[lastTakenIdx++]
		}
		crrMinColumn++
		if lastTakenIdx >= len(ciphertext) {
			break
		}

		// Fill Last row
		for i := crrMinColumn; i <= crrMaxColumn; i++ {
			grid[crrMaxRow] =
				replaceAtIndex(grid[crrMaxRow], ciphertext[lastTakenIdx], i)
			lastTakenIdx++
		}
		crrMaxRow--
		if lastTakenIdx >= len(ciphertext) {
			break
		}
	}
}

func replaceAtIndex(source string, replacement uint8, idx int) string {
	replaced := []rune(source)
	replaced[idx] = rune(replacement)
	return string(replaced)
}
