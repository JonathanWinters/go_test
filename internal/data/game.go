package data

const (
	OPEN_TILE                uint = 0
	WALL                     uint = 1
	PIT_TRAP                 uint = 2
	ARROW_TRAP               uint = 3
	PLAYER_STARTING_POSITION uint = 4
	MAX_DIMENSION            int  = 100
)

type Row []uint
type Map []Row
type Positon struct {
	X int
	Y int
}
