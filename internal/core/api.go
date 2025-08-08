package core

import (
	"github.com/JonathanWinters/go_test/internal/data"
	"github.com/JonathanWinters/go_test/internal/definitions"
)

// !FIXED
// !INFO EFC: struct tags are pieces of metadata attached to the fields of a struct.
// !INFO EFC: they provide instructions or additional information to other Go code or libraries that process the struct
// !INFO EFC: use cases: serialization/deserialization, required fields, empty fields
type SubmitRequest struct {
	UserID definitions.UserID `json:"userid"`
	Level  data.Map           `json:"level"`
}

type SubmitResponse struct {
	Error      string              `json:"error"`
	PrimaryKey int                 `json:"primarykey"`
	LevelID    definitions.LevelID `json:"levelid"`
	Map        data.Map            `json:"map"`
	Position   data.Positon        `json:"position"`
}

type MoveRequest struct {
	PrimaryKey int  `json:"primarykey"`
	Move       int  `json:"move"`
	GodMode    bool `json:"godmode"`
}

type MoveResponse struct {
	Error           string       `json:"error"`
	Result          string       `json:"result"`
	PlayerHitPoints int          `json:"playerhitpoints"`
	Position        data.Positon `json:"position"`
	LatestMap       data.Map     `json:"latestmap"`
}
