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
	if invalid {
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

	allowed, trapHit, oob, result := nextMoveAllowed(newPos, level)
	if allowed || (!oob && moveRequest.GodMode) {
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

func nextMoveAllowed(newPos data.Positon, level data.Map) (allowed bool, trapHit bool, oob bool, result string) {
	maxXIndex := len(level) - 1
	maxYIndex := len(level[0]) - 1

	x := newPos.X
	y := newPos.Y

	if x > maxXIndex || x < 0 {
		allowed = false
		trapHit = false
		oob = true
		result = "X Out of Bounds"
		return
	}

	if y > maxYIndex || y < 0 {
		allowed = false
		trapHit = false
		oob = true
		result = "Y Out of Bounds"
		return
	}

	moveType := level[y][x]

	switch moveType {
	case data.PIT_TRAP, data.ARROW_TRAP:
		allowed = true
		trapHit = true
		oob = false
		return
	case data.WALL:
		allowed = false
		trapHit = false
		oob = false
		result = "Wall Hit"
		return
	}

	allowed = true
	trapHit = false
	oob = false
	return
}

func invalidMove(move int) (err string, invalid bool) {
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
