package core

import (
	"github.com/JonathanWinters/go_test/internal/data"
	"github.com/JonathanWinters/go_test/internal/definitions"
	request "github.com/JonathanWinters/go_test/internal/util/types"
)

// !INFO EFC: struct tags are pieces of metadata attached to the fields of a struct.
// !INFO EFC: they provide instructions or additional information to other Go code or libraries that process the struct
// !INFO EFC: use cases: serialization/deserialization, required fields, empty fields
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
	GodMode    bool
}

type MoveResponse struct {
	Error           string
	Result          string
	PlayerHitPoints int
	Position        data.Positon
	LatestMap       data.Map
}
