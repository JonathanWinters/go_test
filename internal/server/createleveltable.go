package server

import (
	"fmt"
	"net/http"

	"github.com/JonathanWinters/go_test/internal/database"
)

func HandleCreateLevelTable(w http.ResponseWriter, r *http.Request) {
	err := database.CreateLevelTable()
	if err != nil {
		fmt.Fprintf(w, "Level Table Was NOT Created!")
		return
	}
	fmt.Fprintf(w, "Level Table Created!")
}
