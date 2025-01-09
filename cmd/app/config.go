package main

import (
	"net/http"

	"github.com/umeh-promise/blog/internal/controller/middlewares"
	"github.com/umeh-promise/blog/internal/utils"
	"go.uber.org/zap"
)

type application struct {
	config       baseConfig
	logger       *zap.SugaredLogger
	rateLimitter middlewares.RateLimitter
}

type baseConfig struct {
	addr         string
	env          string
	db           dbConfig
	rateLimitter middlewares.RateLimitterConfig
}

type dbConfig struct {
	addr        string
	maxOpenConn int
	maxIdleConn int
	maxIdleTime string
}

func (app *application) RateLimitMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if allow, retryAfter := app.rateLimitter.Allow(r.RemoteAddr); !allow {
			utils.RateLimitExceededResponse(w, r, retryAfter.String())
			return
		}

		next.ServeHTTP(w, r)
	})
}
