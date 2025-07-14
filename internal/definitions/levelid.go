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

// create new one
func NewLevelIDPtr() *LevelID {
	return NewLevelID().AsPtr()
}

// return as a pointer or nil
func (id LevelID) AsPtr() *LevelID {
	if id.IsNil() {
		return nil
	}
	return &id
}

// safely dereference
func (ptr *LevelID) Self() (id LevelID) {
	if ptr != nil {
		id = *ptr
	}
	return
}
