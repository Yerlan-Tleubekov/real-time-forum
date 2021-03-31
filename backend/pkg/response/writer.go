package response

import (
	"encoding/json"
	"net/http"
)

func WriteResponse(w http.ResponseWriter, code int, resp interface{}) {
	w.WriteHeader(code)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
