package util

import "github.com/JonathanWinters/go_test/internal/data"

func FindIndex2DArray(matrix []data.Row, targetValue uint) data.Positon {

	found := false
	rowIndex := -1
	colIndex := -1

	for r, row := range matrix {
		// Iterate through columns in each row
		for c, value := range row {
			if value == targetValue {
				found = true
				rowIndex = r
				colIndex = c
				break // Exit inner loop once found
			}
		}
		if found {
			break // Exit outer loop once found
		}
	}

	pos := data.Positon{X: rowIndex, Y: colIndex}

	return pos
}
