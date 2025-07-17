package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/JonathanWinters/go_test/internal/core"
	"github.com/JonathanWinters/go_test/internal/data"
	"github.com/JonathanWinters/go_test/internal/definitions"
	"github.com/JonathanWinters/go_test/internal/util/types"
)

func HandleSubmit(w http.ResponseWriter, r *http.Request) {
	// Get a single query parameter
	level := r.URL.Query().Get("level")
	fmt.Fprintf(w, "level, %s!\n", level)

	var levelParsed []data.Row
	if err := json.Unmarshal([]byte(level), &levelParsed); err != nil {
		panic(err)
	}

	useridstring := r.URL.Query().Get("userid")
	userid := definitions.UserIDFromString(useridstring)

	submitRequest := core.SubmitRequest{
		RequestType: types.PUT,
		UserID:      userid,
		Level:       levelParsed,
	}

	core.HandleSubmit(w, submitRequest)
}
