package handlers

import (
	"encoding/json"
	"kevinPicon/go/src/CvPro/server"
	"net/http"
)

type HomeResponse struct {
	Message string `json:"message"`
	Status  bool   `json:"status"`
}

func HomeHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		response := HomeResponse{
			Message: "Welcome to Cv Pro",
			Status:  true,
		}
		json.NewEncoder(w).Encode(response)
	}
}
