package core

import (
	"github.com/JonathanWinters/go_test/internal/data"
	"github.com/JonathanWinters/go_test/internal/definitions"
	request "github.com/JonathanWinters/go_test/internal/util/types"
)

type User struct {
	UserID         definitions.UserID
	CurrentLevelID *definitions.LevelID
}

type SubmitRequest struct {
	RequestType request.HttpMethod
	UserID      definitions.UserID
	Level       data.Map
}

type SubmitResponse struct {
	Error      string
	PrimaryKey int
	LevelID    definitions.LevelID
	Map        data.Map
	Position   data.Positon
}

type MoveRequest struct {
	PrimaryKey int
	Move       int
}

type MoveResponse struct {
	Error           string
	Result          string
	Map             data.Map
	PlayerHitPoints int
	Position        data.Positon
	LatestMap       data.Map
}
