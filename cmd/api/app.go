package main

import (
	"log"
	"net/http"
	"time"

	"github.com/ashenkavinda/go_social_app/internel/config"
	"github.com/ashenkavinda/go_social_app/internel/handlers"
	"github.com/ashenkavinda/go_social_app/internel/repository"
	"github.com/ashenkavinda/go_social_app/internel/service"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Application struct {
	Config config.Config
}

func (app *Application) Mount() http.Handler {

	r := chi.NewRouter()

	// A good base middleware stack
	r.Use(middleware.RequestID)
	r.Use(middleware.ClientIPFromRemoteAddr) // pick one ClientIPFrom* based on your infra, see below
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Set a timeout value on the request context (ctx), that will signal
	// through ctx.Done() that the request has timed out and further
	// processing should be stopped.
	r.Use(middleware.Timeout(60 * time.Second))

	postRepository := repository.NewPostRepository(app.Config.Server.DB)
	userRepository := repository.NewUserRepository(app.Config.Server.DB)
	followerRepository := repository.NewFolloerRepository(app.Config.Server.DB)

	postService := service.NewPostService(postRepository)
	userService := service.NewUserService(userRepository, followerRepository)

	healthHandler := handlers.NewHealthHandler(app.Config.App)
	postHandler := handlers.NewPostHandler(&postService)
	userHandler := handlers.NewUserHandler(&userService)

	r.Route("/v1", func(r chi.Router) {
		r.Get("/health", healthHandler.HealthCheckHandler)

		// post api
		r.Route("/post", func(r chi.Router) {
			r.Post("/", postHandler.Create)
			r.Get("/", postHandler.GetAll)
			r.Get("/{id}", postHandler.GetByID)
			r.Put("/{id}", postHandler.UpdateByID)
			r.Delete("/{id}", postHandler.DeleteByID)
		})

		// user api
		r.Route("/user", func(r chi.Router) {
			r.Post("/", userHandler.Create)
			r.Get("/", userHandler.GetAll)
			r.Get("/{id}", userHandler.GetByID)
			r.Get("/{id}/feed", userHandler.Feed)
			r.Post("/follow", userHandler.Follow)
			r.Post("/unfollow", userHandler.Unfollow)
			r.Put("/{id}", userHandler.UpdateByID)
			r.Delete("/{id}", userHandler.DeleteByID)
		})

	})

	return r
}

func (app *Application) Run(mux http.Handler) error {

	ser := http.Server{
		Addr:         app.Config.Server.Addr,
		Handler:      mux,
		WriteTimeout: time.Second * 30,
		ReadTimeout:  time.Second * 10,
		IdleTimeout:  time.Minute,
	}

	log.Printf("Application running on port %s", app.Config.Server.Addr)

	return ser.ListenAndServe()
}
