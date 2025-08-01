package core

import (
	"encoding/json"
	"net/http"

	"github.com/JonathanWinters/go_test/internal/data"
	"github.com/JonathanWinters/go_test/internal/database"
	"github.com/JonathanWinters/go_test/internal/util"
)

func HandleMove(writer http.ResponseWriter, moveRequest MoveRequest) MoveResponse {

	moveResponse := MoveResponse{
		Error:           "",
		Result:          "",
		PlayerHitPoints: 0,
		Position:        data.Positon{X: 0, Y: 0},
		LatestMap:       [][]int{{}},
	}

	err, valid := ValidateMove(moveRequest.Move)

	if !valid {
		moveResponse.Error = err
		return moveResponse
	}

	marhsalledLevel, dbMapErr := database.GetMapByPrimaryKey(moveRequest.PrimaryKey)
	if dbMapErr != nil {
		moveResponse.Error = dbMapErr.Error()
		return moveResponse
	}

	dbPlayerHitPoints, dbHPErr := database.GetPlayerHitPointsByPrimaryKey(moveRequest.PrimaryKey)
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
	dbPosition, dbPosErr := database.GetPositionByPrimaryKey(moveRequest.PrimaryKey)
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

	moveResponse.PlayerHitPoints = dbPlayerHitPoints

	moveResponse.Position = data.Positon{
		X: currentPos.X,
		Y: currentPos.Y,
	}

	copiedMap := make([][]int, len(level))
	for i := range level {
		copiedMap[i] = make([]int, len(level[i]))
		copy(copiedMap[i], level[i]) // Copy elements of inner slice
	}

	moveResponse.LatestMap = copiedMap

	newPos := currentPos
	// Move that current position
	switch moveRequest.Move {
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
	//!INFO EFC: ideally this is where the cheat logic would exist
	allowed, trapHit, ooo, result := NextMoveAllowed(newPos, level)
	if allowed || (!ooo && moveRequest.GodMode) {
		moveResponse.Position = newPos
		moveResponse.Result = "Move Successful"

		if trapHit && !moveRequest.GodMode {
			moveResponse.PlayerHitPoints--
			moveResponse.Result = "Move Successful, Hit Trap"

			if moveResponse.PlayerHitPoints <= 0 { // RESET if hitpoints hits 0
				moveResponse.PlayerHitPoints = 4
				moveResponse.Position = util.FindIndex2DArray(level, 4)
				moveResponse.Result = "Move Successful, Hit Trap and Died, Position Reset, Health Reset"
			}
		}

		marshalledPosition, _ := json.Marshal(moveResponse.Position)

		updateErr := database.UpdateLevelHPAndPositionByPrimaryKey(moveRequest.PrimaryKey, moveResponse.PlayerHitPoints, marshalledPosition)

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
	// testmoveResponse := MoveResponse{
	// 	Error:           "",
	// 	Result:          "",
	// 	PlayerHitPoints: 0,
	// 	Position:        data.Positon{X: 0, Y: 0},
	// 	LatestMap:       [][]int{{}},
	// }
	// return testmoveResponse
}

// !INFO EFC: private funcs (not used outside this package) should always be lowercase (golang auto-enforces this way)
func NextMoveAllowed(newPos data.Positon, level data.Map) (allowed bool, trapHit bool, ooo bool, result string) {
	maxXIndex := len(level) - 1
	maxYIndex := len(level[0]) - 1

	x := newPos.X
	y := newPos.Y

	if x > maxXIndex || x < 0 {
		allowed = false
		trapHit = false
		ooo = true
		result = "X Out of Bounds"
		return
	}

	if y > maxYIndex || y < 0 {
		allowed = false
		trapHit = false
		ooo = true
		result = "Y Out of Bounds"
		return
	}

	moveType := level[y][x]

	switch moveType {
	case data.PIT_TRAP, data.ARROW_TRAP:
		allowed = true
		trapHit = true
		ooo = false
		return
	case data.WALL:
		allowed = false
		trapHit = false
		ooo = false
		result = "Wall Hit"
		return
	}

	allowed = true
	trapHit = false
	ooo = false
	return
}

func ValidateMove(move int) (err string, valid bool) {
	//!INFO EFC: no need to pre-define when it's a return value (golang automagically does this using default value for the datatype)
	//!INFO EFC: caveat: types that default to pointers will be nil
	err = ""
	valid = true

	//!INFO EFC: multi-case switch statements, return instead of break (we love short-circuiting in golang)
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
