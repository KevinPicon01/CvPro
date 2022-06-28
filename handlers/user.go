package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/segmentio/ksuid"
	"golang.org/x/crypto/bcrypt"
	"kevinPicon/go/src/CvPro/models"
	"kevinPicon/go/src/CvPro/repository"
	"kevinPicon/go/src/CvPro/server"
	"net/http"
)

const (
	HASH_COST = 8
)

type VerifyUserRequest struct {
	Name string `json:"name"`
}
type SignUpRequest struct {
	Name        string `json:"name"`
	Username    string `json:"username"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	Description string `json:"description"`
	linkedin    string `json:"linkedin"`
	github      string `json:"github"`
	twitter     string `json:"twitter"`
}

func VerifyUserHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var request = VerifyUserRequest{}
		err := json.NewDecoder(r.Body).Decode(&request)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		exist, err := repository.GetUserByName(r.Context(), request.Name)

		if err != nil {
			http.Error(w, "Error GUE", http.StatusInternalServerError)
			return
		}
		if !exist {
			http.Error(w, "Usuario ya existe", http.StatusBadRequest)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(models.Response{
			Status:  true,
			Message: "User available",
		})
	}
}
func SignUpHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("SignUpHandler")
		var request = SignUpRequest{}
		err := json.NewDecoder(r.Body).Decode(&request)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		exist, err := repository.GetUserByName(r.Context(), request.Username)
		if err != nil {
			http.Error(w, "Error GUE", http.StatusInternalServerError)
			return
		}
		if !exist {
			http.Error(w, "Usuario ya existe", http.StatusBadRequest)
			return
		}
		hashedPass, err := bcrypt.GenerateFromPassword([]byte(request.Password), HASH_COST)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		id, err := ksuid.NewRandom()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		var user = models.User{
			Email:       request.Email,
			Name:        request.Name,
			Id:          id.String(),
			Password:    string(hashedPass),
			Username:    request.Username,
			Description: request.Description,
			Linkedin:    request.linkedin,
			Github:      request.github,
			Twitter:     request.twitter,
		}
		err = repository.InsertUser(r.Context(), &user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(models.Response{
			Message: "User created successfully",
		})
	}

}
