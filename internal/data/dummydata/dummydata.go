package dummydata

import (
	"github.com/JonathanWinters/go_test/internal/data"
	"github.com/JonathanWinters/go_test/internal/definitions"
)

var UserID definitions.UserID = definitions.NewUserID()
var LevelID definitions.LevelID = definitions.NewLevelID()

var Map = data.Map{
	{1, 1, 1, 1, 0, 1, 1, 1},
	{1, 0, 0, 0, 0, 0, 0, 1},
	{1, 0, 1, 1, 1, 3, 1, 1},
	{1, 0, 0, 0, 1, 0, 2, 1},
	{1, 1, 1, 0, 1, 1, 0, 1},
	{1, 0, 0, 0, 1, 0, 0, 1},
	{1, 0, 1, 1, 1, 0, 1, 1},
	{1, 0, 0, 4, 0, 0, 0, 1},
	{1, 1, 1, 1, 1, 1, 1, 1},
}

const PlayerHitPoints = 4
