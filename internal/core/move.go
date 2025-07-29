package core

import (
	"encoding/json"
	"net/http"

	"github.com/JonathanWinters/go_test/internal/data"
	"github.com/JonathanWinters/go_test/internal/database"
	"github.com/JonathanWinters/go_test/internal/util"
)

func HandleMove(writer http.ResponseWriter, submitRequest MoveRequest) MoveResponse {

	moveResponse := MoveResponse{
		Error:           "",
		Result:          "",
		Map:             [][]int{{}}, // TODO: Replace OBJ
		PlayerHitPoints: 0,
		Position:        data.Positon{X: 0, Y: 0},
		LatestMap:       [][]int{{}},
	}

	err, valid := ValidateMove(submitRequest.Move)

	if !valid {
		moveResponse.Error = err
		return moveResponse
	}

	marhsalledLevel, dbMapErr := database.GetMapByPrimaryKey(submitRequest.PrimaryKey)
	if dbMapErr != nil {
		moveResponse.Error = dbMapErr.Error()
		return moveResponse
	}

	dbPlayerHitPoints, dbHPErr := database.GetPlayerHitPointsByPrimaryKey(submitRequest.PrimaryKey)
	if dbHPErr != nil {
		moveResponse.Error = dbHPErr.Error()
		return moveResponse
	}

	var level data.Map

	unmarshallLevelErr := json.Unmarshal(marhsalledLevel, &level)
	if unmarshallLevelErr != nil {
		moveResponse.Error = unmarshallLevelErr.Error()
		return moveResponse
	}

	// Find Current Position
	// currentPos := util.FindIndex2DArray(level, 4)
	dbPosition, dbPosErr := database.GetPositionByPrimaryKey(submitRequest.PrimaryKey)
	if dbPosErr != nil {
		moveResponse.Error = dbPosErr.Error()
		return moveResponse
	}

	var currentPos data.Positon

	unmarshallPosErr := json.Unmarshal(dbPosition, &currentPos)

	if unmarshallPosErr != nil {
		moveResponse.Error = unmarshallPosErr.Error()
		return moveResponse
	}

	moveResponse.Map = level
	moveResponse.PlayerHitPoints = dbPlayerHitPoints

	moveResponse.Position = data.Positon{
		X: currentPos.X,
		Y: currentPos.Y,
	}

	copiedMap := make([][]int, len(level))
	for i := range moveResponse.Map {
		copiedMap[i] = make([]int, len(level[i]))
		copy(copiedMap[i], level[i]) // Copy elements of inner slice
	}

	moveResponse.LatestMap = copiedMap

	newPos := currentPos
	// Move that current position
	switch submitRequest.Move {
	case data.MOVE_LEFT:
		newPos.X--
	case data.MOVE_UP:
		newPos.Y--
	case data.MOVE_RIGHT:
		newPos.X++
	case data.MOVE_DOWN:
		newPos.Y++
	}

	//Check if new position exists within the Map
	//Check if new position is going to result in a move, player HP going down
	allowed, trapHit, result := NextMoveAllowed(newPos, level)
	if allowed {
		moveResponse.Position = newPos
		moveResponse.Result = "Move Successful"

		if trapHit {
			moveResponse.PlayerHitPoints--
			moveResponse.Result = "Move Successful, Hit Trap"

			if moveResponse.PlayerHitPoints <= 0 { // RESET if hitpoints hits 0
				moveResponse.PlayerHitPoints = 4
				moveResponse.Position = util.FindIndex2DArray(level, 4)
				moveResponse.Result = "Move Successful, Hit Trap and Died, Position Reset, Health Reset"
			}
		}

		marshalledPosition, _ := json.Marshal(moveResponse.Position)

		updateErr := database.UpdateLevelHPAndPositionByPrimaryKey(submitRequest.PrimaryKey, moveResponse.PlayerHitPoints, marshalledPosition)

		if updateErr != nil {
			moveResponse.Error = "Database Update Failure, " + updateErr.Error()
		}
	} else {
		moveResponse.Result = "Move Not Allowed: " + result
	}

	ogPos := util.FindIndex2DArray(moveResponse.LatestMap, 4)
	moveResponse.LatestMap[moveResponse.Position.Y][moveResponse.Position.X] = data.PLAYER_STARTING_POSITION
	moveResponse.LatestMap[ogPos.Y][ogPos.X] = data.OPEN_TILE

	return moveResponse
}

func NextMoveAllowed(newPos data.Positon, level data.Map) (allowed bool, trapHit bool, result string) {
	maxXIndex := len(level) - 1
	maxYIndex := len(level[0]) - 1

	x := newPos.X
	y := newPos.Y

	if x > maxXIndex || x < 0 {
		allowed = false
		trapHit = false

		result = "X Out of Bounds"
		return
	}

	if y > maxYIndex || y < 0 {
		allowed = false
		trapHit = false

		result = "Y Out of Bounds"
		return
	}

	moveType := level[y][x]

	switch moveType {
	case data.PIT_TRAP, data.ARROW_TRAP:
		allowed = true
		trapHit = true
		return
	case data.WALL:
		allowed = false
		trapHit = false
		result = "Wall Hit"
		return
	}

	allowed = true
	trapHit = false

	return
}

func ValidateMove(move int) (err string, valid bool) {
	err = ""
	valid = true

	switch move {
	case data.MOVE_LEFT:
		break
	case data.MOVE_UP:
		break
	case data.MOVE_RIGHT:
		break
	case data.MOVE_DOWN:
		break
	default:
		err = "Move Value is Invalid"
		valid = false
	}

	return
}
