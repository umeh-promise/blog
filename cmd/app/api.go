package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/umeh-promise/blog/internal/utils"
)

func (app *application) mount(routerGroups ...func(r chi.Router)) *chi.Mux {
	router := chi.NewRouter()

	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{utils.GetString("CORS_ALLOWED_ORIGIN", "https://localhost:4000")},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))
	router.Use(app.RateLimitMiddleware)
	router.Use(middleware.Timeout(60 * time.Second))

	router.Route("/v1", func(router chi.Router) {
		for _, subRouter := range routerGroups {
			router.Group(subRouter)
		}
	})

	return router
}

func (app *application) run(handler *chi.Mux) error {

	server := &http.Server{
		Addr:         app.config.addr,
		Handler:      handler,
		WriteTimeout: time.Second * 30,
		ReadTimeout:  time.Second * 10,
		IdleTimeout:  time.Minute,
	}

	log.Printf("Server has started at %s with environment %s", app.config.addr, app.config.env)

	shutdown := make(chan error)

	go func() {
		quit := make(chan os.Signal, 1)
		s := <-quit

		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

		ctx, cancel := context.WithTimeout(context.Background(), utils.QueryTimeout)
		defer cancel()

		log.Printf("Server signal %s caught", s.String())
		shutdown <- server.Shutdown(ctx)
	}()

	err := server.ListenAndServe()
	if !errors.Is(err, http.ErrServerClosed) {
		return err
	}

	if err = <-shutdown; err != nil {
		return err
	}

	log.Printf("Server existed, addr %s and environment %s", app.config.addr, app.config.env)

	return nil
}
