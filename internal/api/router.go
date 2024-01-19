package api

import (
	v1 "user-app/internal/api/controller/v1"
	"user-app/internal/consts"
	"user-app/internal/middleware"
	"user-app/internal/repository"
	"user-app/internal/service"

	"github.com/go-chi/chi"
)

func (s *Server) InitRouter() chi.Router {
	r := chi.NewRouter()
	r.Use(middleware.LoggingMiddleware)
	r.Mount(consts.BasePath, s.initRouteV1())

	return r
}

func (s *Server) initRouteV1() chi.Router {

	// repo layer
	userRepo := repository.NewUserRepository(s.sq, s.db) // Pass s.db here

	// service layer
	userService := service.NewUserService(userRepo)

	// controllers
	userCtl := v1.NewUserController(userService)

	r := chi.NewRouter()
	r.Route("/users", func(r chi.Router) {
		r.With(middleware.ValidateCreateUserRequest).Post("/", userCtl.CreateUser)
		r.With(middleware.ValidateUpdateUserRequest).Put("/{userID}", userCtl.UpdateUser)
	})

	return r
}
