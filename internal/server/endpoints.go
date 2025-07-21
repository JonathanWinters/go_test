package server

import (
	"net/http"
)

func SetHandlers() {

	http.HandleFunc("/", HandleBase)
	http.HandleFunc("/submit", HandleSubmit)
	http.HandleFunc("/createleveltable", HandleCreateLevelTable)
}
