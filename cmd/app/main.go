package main

import (
	"time"

	_ "github.com/lib/pq"
	"github.com/umeh-promise/blog/internal/controller/handlers"
	"github.com/umeh-promise/blog/internal/controller/middlewares"
	"github.com/umeh-promise/blog/internal/controller/routes"
	"github.com/umeh-promise/blog/internal/db"
	"github.com/umeh-promise/blog/internal/repositories"
	"github.com/umeh-promise/blog/internal/services"
	"github.com/umeh-promise/blog/internal/utils"
	"go.uber.org/zap"
)

func main() {
	config := baseConfig{
		addr: utils.GetString("ADDR", ":8080"),
		env:  utils.GetString("ENV", "development"),
		db: dbConfig{
			addr:        utils.GetString("DB_ADDR", "postgres://user:password@localhost:5432/blog?sslmode=disable"),
			maxOpenConn: utils.GetInt("DB_MAX_OPEN_CONNS", 30),
			maxIdleConn: utils.GetInt("DB_MAX_IDLE_CONNS", 30),
			maxIdleTime: utils.GetString("DB_MAX_IDLE_TIME", "15m"),
		},
		rateLimitter: middlewares.RateLimitterConfig{
			RequestPerTimeFrame: utils.GetInt("RATELIMITER_REQUEST_COUNT", 20),
			TimeFrame:           time.Second * 5,
		},
	}

	logger := zap.Must(zap.NewProduction()).Sugar()
	defer logger.Sync()

	db, err := db.NewDBConnection(config.db.addr, config.db.maxOpenConn, config.db.maxIdleConn, config.db.maxIdleTime)
	if err != nil {
		logger.Fatal("failed to open database connection %w", err)
	}
	defer db.Close()
	logger.Info("DB connected successfully")

	// Users
	userRepo := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService)
	authMiddleware := middlewares.NewAuthMiddleware(userService)

	// Role
	roleRepo := repositories.NewRoleRepository(db)
	roleService := services.NewRoleService(roleRepo)
	roleMiddlware := middlewares.NewRoleMiddleware(roleService)

	// Comment
	commentRepo := repositories.NewCommentRepository(db)
	commentService := services.NewCommentService(commentRepo)
	commentHandler := handlers.NewCommentHandler(commentService)

	// Posts
	postRepo := repositories.NewPostRepository(db)
	postService := services.NewPostService(postRepo)
	postHandler := handlers.NewPostHandler(postService, commentService)
	postMiddleware := middlewares.NewPostMidleware(postService)

	// Rate limiter
	rateLimiter := middlewares.NewFixedWindowLimiter(
		config.rateLimitter.RequestPerTimeFrame, config.rateLimitter.TimeFrame,
	)

	app := &application{
		config:       config,
		logger:       logger,
		rateLimitter: rateLimiter,
	}

	router := app.mount(
		routes.PostRouter(postHandler, commentHandler, postMiddleware, authMiddleware, roleMiddlware),
		routes.UserRouter(userHandler, authMiddleware),
	)
	logger.Fatal(app.run(router))
}
