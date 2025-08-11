package core

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/JonathanWinters/go_test/internal/data"
	"github.com/JonathanWinters/go_test/internal/database"
	"github.com/JonathanWinters/go_test/internal/definitions"
	"github.com/JonathanWinters/go_test/internal/util"
)

type WriteError struct {
	Message     string
	ResultLevel [][]int
	UserId      definitions.UserID
}

func (e *WriteError) Error() string {
	return fmt.Sprintf("Error %s: %d - %s", e.Message, e.ResultLevel, e.UserId)
}

type ValidatationError struct {
	Message string
}

func (e *ValidatationError) Error() string {
	return fmt.Sprintf("Error: %s", e.Message)
}

func HandleSubmit(writer http.ResponseWriter, submitRequest SubmitRequest) (submitResponse SubmitResponse) {

	err := ValidateMapSubmission(submitRequest.Level)
	if err != nil {
		// return what went wrong
		// fmt.Fprintf(writer, "Validity:, %s!\n Invalid")
		var ve *ValidatationError
		errors.As(err, &ve)
		Error := WriteError{
			Message:     ve.Message,
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
			Position:   data.Positon{X: 0, Y: 0},
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
	Position := util.FindIndex2DArray(levelMap, 4)

	levelSubmission := database.Level{
		ID:              levelID,
		Map:             levelMap,
		Position:        Position,
		PlayerHitPoints: 4,
	}

	pk, err := database.InsertLevel(levelSubmission)

	if err != nil {
		log.Fatal(err)
	}

	submitResponse = SubmitResponse{
		PrimaryKey: pk,
		LevelID:    levelID,
		Map:        levelMap,
		Position:   Position,
	}
	return submitResponse
}

// 1. Maps must be retangular
// 2. Maps may not be large than 100 in any dimenion
// 3. Map spaces may not use values other the number 0, 1, 2, 3, or 4.
func ValidateMapSubmission(matrix data.Map) (err error) {

	firstRowLen := len(matrix[0])
	colLen := len(matrix)

	if !ValidateDimensions(colLen) {
		err = &ValidatationError{
			Message: "Dimensions Error: colLen",
		}
		return
	}

	for r, row := range matrix {
		if !ValidateRectangle(firstRowLen, row) {
			GetObfuscatedError(RECTANGULAR)
			err = &ValidatationError{
				Message: "Rectangle Error",
			}
			return
		}

		rowLen := len(matrix[r])
		if !ValidateDimensions(rowLen) {
			err = &ValidatationError{
				Message: "Dimensions Error: rowLen",
			}
			return
		}
		// Iterate through columns in each row
		for _, value := range row {

			if !ValidateMapValues(value) {
				GetObfuscatedError(VALUES)
				err = &ValidatationError{
					Message: "Map Values Error",
				}
				return
			}
		}
	}
	return
}

// !NOTED
// !INFO EFC: appreciate all these checks being broke out into simple, readable funcs
func ValidateRectangle(firstRowLen int, row []int) bool {
	return len(row) == firstRowLen
}

func ValidateDimensions(length int) bool {
	return length <= data.MAX_DIMENSION
}

func ValidateMapValues(value int) bool {
	switch value {
	case data.OPEN_TILE, data.WALL, data.PIT_TRAP, data.ARROW_TRAP, data.PLAYER_STARTING_POSITION:
		return true
	default:
		return false
	}
}
