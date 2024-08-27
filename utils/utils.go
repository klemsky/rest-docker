package utils

import (
	"encoding/json"
	"net/http"
)

func RespondWithError(w http.ResponseWriter, code int, internalCode int, errMsg string) {
	RespondWithJSON(w, code, map[string]interface{}{"code": internalCode, "msg": "Failed", "error": errMsg})
}

func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
