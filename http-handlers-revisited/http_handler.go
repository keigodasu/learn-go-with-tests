package http_handlers_revisited

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	Name string
}

type UserService interface {
	Register(user User) (insertedID string, err error)
}

type UserServer struct {
	service UserService
}

func NewUserServer(service UserService) *UserServer {
	return &UserServer{service: service}
}

func (u *UserServer) RegisterUser(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var newUser User
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		http.Error(w, fmt.Sprintf("could not decode user paload %v", err), http.StatusBadRequest)
		return
	}

	insertedfID, err := u.service.Register(newUser)
	if err != nil {
		http.Error(w, fmt.Sprintf("%v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprint(w, insertedfID)
}
