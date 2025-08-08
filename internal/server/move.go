//go:build !cheats

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

const CheatsAllowed = false

var MoveRequestQueue = []core.MoveRequest{}

func HandleMove(w http.ResponseWriter, r *http.Request) {

	//!FIXED
	// !INFO EFC: all this somewhat complicated marshalling/demarshalling is luckily taken care of us by our RP package :D
	mrb, err := DecodeJson(w, r)
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
