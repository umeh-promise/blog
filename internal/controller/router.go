package controller

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
)

type Server struct {
	Engine *chi.Mux
	Server *http.Server
}

func NewServer(routerMappings func(router *chi.Mux)) *Server {
	router := chi.NewRouter()

	// router mappings
	routerMappings(router)

	server := &http.Server{
		Addr:         "0.0.0.0:8080",
		Handler:      router,
		WriteTimeout: time.Second * 30,
		ReadTimeout:  time.Second * 10,
		IdleTimeout:  time.Minute,
	}

	return &Server{
		Engine: router,
		Server: server,
	}
}

func (s *Server) Serve() error {

	shutdown := make(chan error)

	go func() {
		quit := make(chan os.Signal, 1)
		sig := <-quit

		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		log.Printf("Signal caught %s", sig.String())

		shutdown <- s.Server.Shutdown(ctx)
	}()

	if err := s.Server.ListenAndServe(); errors.Is(err, http.ErrServerClosed) {
		return err
	}

	if err := <-shutdown; err != nil {
		return err
	}

	log.Println("Server existed properly")

	return nil
}
