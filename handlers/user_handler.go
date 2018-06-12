package handlers

import (
	"fmt"
	"github.com/dima-kov/go-architecture/services"
	"goji.io/pat"
	"net/http"
	"strconv"
)

type UserHandler struct {
	service services.UserService
}

func NewUserHandler(service services.UserService) UserHandler {
	return UserHandler{service: service}
}

func (h UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	userIDStr := pat.Param(r, "user-id")
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("User-id is not a integer"))
		return
	}

	user, err := h.service.GetUser(uint(userID))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
	fmt.Println(user)
	// return user info here
}
