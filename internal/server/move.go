//go:build !cheats

package server

import (
	"log"
	"net/http"

	"github.com/JonathanWinters/go_test/internal/core"
)

const CheatsAllowed = false

var MoveRequestQueue = []core.MoveRequest{}

func HandleMove(w http.ResponseWriter, r *http.Request) {

	mrb, err := DecodeMoveJson(w, r)
	if err != nil {
		return
	}

	moveRequest := core.MoveRequest{
		PrimaryKey: mrb.PrimaryKey,
		Move:       mrb.Move,
		GodMode:    CheatsAllowed,
	}

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
