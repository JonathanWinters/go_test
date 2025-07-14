package definitions

import uuid "github.com/JonathanWinters/go_test/internal/util/types"

// ID of a specific wager
type UserID struct {
	uuid.UUID
}

// create new one
func NewUserID() (id UserID) {
	id.UUID = uuid.New()
	return
}
