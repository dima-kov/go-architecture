package server

import (
	"github.com/dima-kov/go-architecture/handlers"
	"github.com/dima-kov/go-architecture/repositories"
	"github.com/dima-kov/go-architecture/server/router"
	"github.com/dima-kov/go-architecture/services"
	"github.com/jinzhu/gorm"
	"net/http"
)

type Server struct {
	routes router.Router
}

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

	s.routes = router.NewRouter()
	s.routes.CreateRoutes(userHandler, postHandler)

	serve := &http.Server{Addr: ":8000", Handler: s.routes.Route}
	serve.ListenAndServe()
}
