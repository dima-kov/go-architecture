package handlers

import (
	"fmt"
	"github.com/dima-kov/go-architecture/services"
	"goji.io/pat"
	"net/http"
	"strconv"
)

type PostHandler struct {
	service services.PostService
}

func NewPostHandler(service services.PostService) PostHandler {
	return PostHandler{service: service}
}

func (h PostHandler) GetPost(w http.ResponseWriter, r *http.Request) {
	postIDStr := pat.Param(r, "post-id")
	postID, err := strconv.Atoi(postIDStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Post-id is not an integer"))
		return
	}

	post, err := h.service.GetPost(uint(postID))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
	fmt.Println(post)
	// return post info here
}
