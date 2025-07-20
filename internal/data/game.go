package data

const (
	OPEN_TILE                int = 0
	WALL                     int = 1
	PIT_TRAP                 int = 2
	ARROW_TRAP               int = 3
	PLAYER_STARTING_POSITION int = 4
	MAX_DIMENSION            int = 100
)

type Map [][]int
type Positon struct {
	X int
	Y int
}
