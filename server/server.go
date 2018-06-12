package server

import (
	"github.com/dima-kov/go-architecture/handlers"
	"github.com/dima-kov/go-architecture/repositories"
	"github.com/dima-kov/go-architecture/services"
	"github.com/jinzhu/gorm"
	"goji.io"
	"goji.io/pat"
	"net/http"
)

type Server struct{}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) Start() {
	dbConnection, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer dbConnection.Close()

	userRepository := repositories.NewUserRepository(dbConnection)
	postRepository := repositories.NewPostRepository(dbConnection)

	userService := services.NewUserService(userRepository)
	postService := services.NewPostService(postRepository)

	userHandler := handlers.NewUserHandler(userService)
	postHandler := handlers.NewPostHandler(postService)

	route := goji.NewMux()
	route.HandleFunc(
		pat.Get("/user/:user-id/"),
		userHandler.GetUser,
	)
	route.HandleFunc(
		pat.Post("/post/:post-id/"),
		postHandler.GetPost,
	)

	server := &http.Server{Addr: ":8000", Handler: route}
	server.ListenAndServe()
}
