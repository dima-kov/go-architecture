package router

import (
	"github.com/dima-kov/go-architecture/handlers"
	"goji.io"
	"goji.io/pat"
)

type Router struct {
	Route *goji.Mux
}

func NewRouter() Router {
	return Router{Route: goji.NewMux()}
}

func (r *Router) CreateRoutes(userHandler handlers.UserHandler, postHandler handlers.PostHandler) {
	r.Route.HandleFunc(
		pat.Get("/user/:user-id/"),
		userHandler.GetUser,
	)
	r.Route.HandleFunc(
		pat.Post("/post/:post-id/"),
		postHandler.GetPost,
	)
}
