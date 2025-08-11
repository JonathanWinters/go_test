//go:build windows

package server

import (
	"log"
	"net/http"
)

func HandleSubmit(w http.ResponseWriter, r *http.Request) {

	rawResult, err := ProcessSubmitRequest(w, r)
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
