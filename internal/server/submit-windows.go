//go:build windows

package server

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/JonathanWinters/go_test/internal/core"
	"github.com/JonathanWinters/go_test/internal/definitions"
	"github.com/JonathanWinters/go_test/internal/util"
)

type RequestBody struct {
	UserId string
	Level  [][]int
}

func HandleSubmit(w http.ResponseWriter, r *http.Request) {

	notJson := util.HandleContentTypeError(w, r)
	if notJson {
		return
	}

	r.Body = http.MaxBytesReader(w, r.Body, util.MaxMB)

	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()

	var p RequestBody
	err := dec.Decode(&p)
	if err != nil {
		util.HandleDecodeError(w, err)
		return
	}

	responseObj := RequestBody{
		UserId: p.UserId,
		Level:  p.Level,
	}

	userid := definitions.UserIDFromString(responseObj.UserId)

	submitRequest := core.SubmitRequest{
		UserID: userid,
		Level:  responseObj.Level,
	}

	submitResponse := core.HandleSubmit(w, submitRequest)

	rawResult, err := json.Marshal(submitResponse)
	if err != nil {
		log.Fatal(err)
		return
	}

	//!FIXED
	// !INFO EFC: result should be returned as a response
	// fmt.Fprintf(w, "%s", rawResult)
	_, err = w.Write(rawResult)
	if err != nil {
		log.Fatal(err)
		return
	}

}
