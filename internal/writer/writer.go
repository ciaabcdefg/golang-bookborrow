package writer

import (
	"encoding/json"
	"net/http"
)

func Write(w http.ResponseWriter, body string, status int) {
	w.WriteHeader(status)
	w.Write([]byte(body))
}

func WriteOk(w http.ResponseWriter, msg string, status int) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(map[string]string{
		"message": msg,
	})
}

func WriteError(w http.ResponseWriter, err error, status int) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(map[string]string{
		"error": err.Error(),
	})
}

func WriteJson(w http.ResponseWriter, v any, status int) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(v)
}
