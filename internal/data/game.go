package data

// !INFO EFC: I prefer to be explicit, but golang does let you exclude types after defining the first
const (
	OPEN_TILE                int = 0
	WALL                     int = 1
	PIT_TRAP                 int = 2
	ARROW_TRAP               int = 3
	PLAYER_STARTING_POSITION int = 4
	MAX_DIMENSION            int = 100
)

const (
	MOVE_LEFT  int = 0
	MOVE_UP    int = 1
	MOVE_RIGHT int = 2
	MOVE_DOWN  int = 3
)

type Map [][]int
type Positon struct {
	X int
	Y int
}

const DBConnectionString = "postgres://postgres:secret@localhost:5434/pg-test?sslmode=disable"
