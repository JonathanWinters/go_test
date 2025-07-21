package core

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/JonathanWinters/go_test/internal/data"
	"github.com/JonathanWinters/go_test/internal/database"
	"github.com/JonathanWinters/go_test/internal/definitions"
	"github.com/JonathanWinters/go_test/internal/util"
)

type WriteError struct {
	Error       string
	ResultLevel [][]int
	UserId      definitions.UserID
}

type ValidatationError struct {
	Error string
}

func HandleSubmit(writer http.ResponseWriter, submitRequest SubmitRequest) SubmitResponse {

	validationError, valid := ValidateMapSubmission(submitRequest.Level)
	if !valid {
		// return what went wrong
		// fmt.Fprintf(writer, "Validity:, %s!\n Invalid")
		Error := WriteError{
			Error:       validationError.Error,
			ResultLevel: submitRequest.Level,
			UserId:      submitRequest.UserID,
		}

		emptyMap := [][]int{{}}
		errorLevelID := definitions.NewLevelID()

		submitResponse := SubmitResponse{
			Error:      "Validation Error",
			PrimaryKey: 0,
			LevelID:    errorLevelID,
			Map:        emptyMap,
		}

		rawError, err := json.Marshal(Error)
		if err != nil {
			return submitResponse
		}

		fmt.Fprintf(writer, "%s", rawError)
		return submitResponse
	}

	levelID := definitions.NewLevelID()
	levelMap := submitRequest.Level
	originalPosition := util.FindIndex2DArray(levelMap, 4)

	levelSubmission := database.Level{
		ID:               levelID,
		Map:              levelMap,
		OriginalPosition: originalPosition,
		PlayerHitPoints:  4,
	}

	pk := database.InsertLevel(levelSubmission)

	submitResponse := SubmitResponse{
		PrimaryKey: pk,
		LevelID:    levelID,
		Map:        levelMap,
	}
	return submitResponse
}

// 1. Maps must be retangular
// 2. Maps may not be large than 100 in any dimenion
// 3. Map spaces may not use values other the number 0, 1, 2, 3, or 4.
func ValidateMapSubmission(matrix data.Map) (validateError ValidatationError, valid bool) {

	firstRowLen := len(matrix[0])
	colLen := len(matrix)

	if !ValidateDimensions(colLen) {
		valid = false
		validateError = ValidatationError{
			Error: "Dimensions Error: colLen",
		}
		return
	}

	for r, row := range matrix {
		if !ValidateRectangle(firstRowLen, row) {
			GetObfuscatedError(RECTANGULAR)
			valid = false
			validateError = ValidatationError{
				Error: "Rectangle Error",
			}
			return
		}

		rowLen := len(matrix[r])
		if !ValidateDimensions(rowLen) {
			valid = false
			validateError = ValidatationError{
				Error: "Dimensions Error: rowLen",
			}
			return
		}
		// Iterate through columns in each row
		for _, value := range row {

			if !ValidateMapValues(value) {
				GetObfuscatedError(VALUES)
				valid = false
				validateError = ValidatationError{
					Error: "Map Values Error",
				}
				return
			}
		}
	}
	valid = true
	validateError = ValidatationError{
		Error: "N/A",
	}
	return
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
