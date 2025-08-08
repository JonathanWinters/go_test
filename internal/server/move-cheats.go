//go:build cheats

package server

import (
	"log"
	"net/http"

	"github.com/JonathanWinters/go_test/internal/core"
)

type MoveRequestBody struct {
	PrimaryKey int
	Move       int
	GodMode    bool
}

var MoveRequestQueue = []core.MoveRequest{}

// !FIXED
// !INFO EFC: for build flag code, try to isolate only the meaningful differences, reduce dupe code
func HandleMove(w http.ResponseWriter, r *http.Request) {

	mrb, err := DecodeJson(w, r)
	if err != nil {
		return
	}

	moveRequest := core.MoveRequest{
		PrimaryKey: mrb.PrimaryKey,
		Move:       mrb.Move,
		GodMode:    mrb.GodMode,
	}

	log.Println("Decode Successful")

	rawResult, err := ProcessMoveRequest(w, moveRequest)
	if err != nil {
		log.Fatal(err)
		return
	}

	_, err = w.Write(rawResult)
	if err != nil {
		log.Fatal(err)
		return
	}
}
