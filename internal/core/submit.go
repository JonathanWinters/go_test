package core

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/JonathanWinters/go_test/internal/data"
	"github.com/JonathanWinters/go_test/internal/definitions"
)

type WriteError struct {
	Error       string
	ResultLevel [][]int
	UserId      definitions.UserID
}

func HandleSubmit(writer http.ResponseWriter, submitRequest SubmitRequest) { //(submitResponse SubmitResponse, err error)

	if !ValidateMapSubmission(submitRequest.Level) {
		// return what went wrong
		// fmt.Fprintf(writer, "Validity:, %s!\n Invalid")
		Error := WriteError{
			Error:       "Shit is fucked mate",
			ResultLevel: submitRequest.Level,
			UserId:      submitRequest.UserID,
		}

		rawError, err := json.Marshal(Error)
		if err != nil {
			return
		}

		fmt.Fprintf(writer, "%s", rawError)
		return
	}

	rawResult, err := json.Marshal(submitRequest)
	if err != nil {
		return
	}

	fmt.Fprintf(writer, "%s", rawResult)
	// fmt.Fprintf(writer, "Validity:, %s!\n VALID")
}

// 1. Maps must be retangular
// 2. Maps may not be large than 100 in any dimenion
// 3. Map spaces may not use values other the number 0, 1, 2, 3, or 4.
func ValidateMapSubmission(matrix data.Map) bool {

	firstRowLen := len(matrix[0])
	colLen := len(matrix)

	if ValidateDimensions(colLen) {
		return false
	}

	for r, row := range matrix {
		if !ValidateRectangle(firstRowLen, row) {
			GetObfuscatedError(RECTANGULAR)
			return false
		}

		rowLen := len(matrix[r])
		if !ValidateDimensions(rowLen) {
			return false
		}
		// Iterate through columns in each row
		for _, value := range row {

			if !ValidateMapValues(value) {
				GetObfuscatedError(VALUES)
				return false
			}
		}
	}
	return true
}

func ValidateRectangle(firstRowLen int, row []int) bool {
	return len(row) == firstRowLen
}

func ValidateDimensions(length int) bool {
	return length <= data.MAX_DIMENSION
}

func ValidateMapValues(value int) bool {
	switch value {
	case data.OPEN_TILE:
	case data.WALL:
	case data.PIT_TRAP:
	case data.ARROW_TRAP:
	case data.PLAYER_STARTING_POSITION:
		return true
	default:
		return false
	}
	return true
}
