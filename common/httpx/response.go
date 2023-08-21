package httpx

import (
	"encoding/json"
	"net/http"
)

func JSONResponse(w http.ResponseWriter, response interface{}, code int) {
	var data []byte
	var err error

	val, ok := response.([]byte)
	if ok {
		data = val
	} else {
		data, err = json.Marshal(response)
	}

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error())) // nolint:errcheck
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(data) // nolint:errcheck
}
