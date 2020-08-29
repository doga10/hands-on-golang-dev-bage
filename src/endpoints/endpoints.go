package endpoints

import (
	"encoding/json"
	"net/http"
)

type EndpointResponse struct {
	Status bool        `json:"status"`
	Code   int         `json:"code"`
	Data   interface{} `json:"data"`
}

func Status(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	response := EndpointResponse{Status: true, Code: 200, Data: "Status OK"}
	json.NewEncoder(w).Encode(response)
}
