package core

import (
	"encoding/json"
	"fmt"
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

	err, invalid := invalidMove(moveRequest.Move)
	moveResponse.Error = "BEFORE invalid Check"
	if invalid {
		moveResponse.Error = err
		return moveResponse
	}
	moveResponse.Error = "AFTER invalid Check"
	marhsalledLevel, dbMapErr := database.GetMapByPrimaryKey(moveRequest.PrimaryKey)
	moveResponse.Error = "BEFORE dbMapErr Check"
	if dbMapErr != nil {
		moveResponse.Error = dbMapErr.Error()
		return moveResponse
	}
	moveResponse.Error = "AFTER dbMapErr Check"
	dbPlayerHitPoints, dbHPErr := database.GetPlayerHitPointsByPrimaryKey(moveRequest.PrimaryKey)
	moveResponse.Error = "BEFORE dbHPErr Check"
	if dbHPErr != nil {
		moveResponse.Error = dbHPErr.Error()
		return moveResponse
	}
	moveResponse.Error = "AFTER dbHPErr Check"
	var level data.Map

	unmarshallLevelErr := json.Unmarshal(marhsalledLevel, &level)
	moveResponse.Error = "BEFORE unmarshallLevelErr Check"
	if unmarshallLevelErr != nil {
		moveResponse.Error = unmarshallLevelErr.Error()
		return moveResponse
	}
	moveResponse.Error = "AFTER unmarshallLevelErr Check"
	// Find Current Position
	dbPosition, dbPosErr := database.GetPositionByPrimaryKey(moveRequest.PrimaryKey)
	moveResponse.Error = "BEFORE dbPosErr Check"
	if dbPosErr != nil {
		moveResponse.Error = dbPosErr.Error()
		return moveResponse
	}
	moveResponse.Error = "AFTER dbPosErr Check"
	var currentPos data.Positon

	unmarshallPosErr := json.Unmarshal(dbPosition, &currentPos)
	moveResponse.Error = "BEFORE unmarshallPosErr Check"
	if unmarshallPosErr != nil {
		moveResponse.Error = unmarshallPosErr.Error()
		return moveResponse
	}
	moveResponse.Error = "AFTER unmarshallPosErr Check"
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
	allowed, trapHit, ooo, result := nextMoveAllowed(newPos, level)
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

	isAtSameY := moveResponse.Position.Y == ogPos.Y
	isAtSameX := moveResponse.Position.X == ogPos.X

	if !(isAtSameY && isAtSameX) {
		moveResponse.LatestMap[ogPos.Y][ogPos.X] = data.OPEN_TILE
	}

	return moveResponse
}

// !FIXED
// !INFO EFC: private funcs (not used outside this package) should always be lowercase (golang auto-enforces this way)
func nextMoveAllowed(newPos data.Positon, level data.Map) (allowed bool, trapHit bool, ooo bool, result string) {
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

func invalidMove(move int) (err string, invalid bool) {
	//!FIXED
	//!INFO EFC: no need to pre-define when it's a return value (golang automagically does this using default value for the datatype)
	//!INFO EFC: caveat: types that default to pointers will be nil

	//!FIXED
	//!INFO EFC: multi-case switch statements, return instead of break (we love short-circuiting in golang)
	switch move {
	case data.MOVE_LEFT, data.MOVE_UP, data.MOVE_RIGHT, data.MOVE_DOWN:
		err = "Move Value is NOT Invalid: " + fmt.Sprint(move)
		invalid = false
		return
	default:
		err = "Move Value is Invalid: " + fmt.Sprint(move)
		invalid = true
		return
	}
}
