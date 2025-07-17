package types

import (
	google "github.com/google/uuid"
)

// elevate a UUID to a struct, making it more convenient to create user types which are a UUID and get all of the standard behaviors for free
type UUID struct {
	google.UUID // embed a google UUID so that we can tell that we are, and anonymous embeds of us, are UUIDs
}

// create new v4 UUID
func New() (id UUID) {
	id.UUID = google.New()
	return
}

// check if nil UUID
func (id UUID) IsNil() bool {
	return id.UUID == google.Nil
}

func FromString(id string) UUID {
	return UUID{google.MustParse(id)}
}
