package server

import (
	"fmt"
	"net/http"
)

func HandleBase(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Connected to Server!")
}
