package handler

import (
	"encoding/json"
	"net/http"
)

func ReadRequest(r *http.Request, result interface{}) {
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(result)
	ErrorPanic(err)
}
