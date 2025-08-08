package definitions

import uuid "github.com/JonathanWinters/go_test/internal/util/types"

// ID of a specific wager
type LevelID struct {
	uuid.UUID
}

// create new one
func NewLevelID() (id LevelID) {
	id.UUID = uuid.New()
	return
}

//!FIXED
// !INFO EFC: unused funcs
