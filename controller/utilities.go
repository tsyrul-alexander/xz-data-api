package controller

import (
	"encoding/json"
	"log"
	"net/http"
)

func SetResponse(w http.ResponseWriter, data interface{}) {
	w.Header().Add("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Println(err.Error())
	}
}

func decodeJsonBody(r *http.Request, obj interface{}) error {
	var jsonDecoder = json.NewDecoder(r.Body)
	return jsonDecoder.Decode(obj)
}
