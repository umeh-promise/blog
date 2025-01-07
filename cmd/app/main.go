package main

import (
	"log"

	_ "github.com/lib/pq"
	"github.com/umeh-promise/blog/internal/controller/handlers"
	"github.com/umeh-promise/blog/internal/controller/routes"
	"github.com/umeh-promise/blog/internal/db"
	"github.com/umeh-promise/blog/internal/repositories"
	"github.com/umeh-promise/blog/internal/services"
	"github.com/umeh-promise/blog/internal/utils"
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
	}

	db, err := db.NewDBConnection(config.db.addr, config.db.maxOpenConn, config.db.maxIdleConn, config.db.maxIdleTime)
	if err != nil {
		log.Fatal("failed to open database connection %w", err)
	}
	defer db.Close()
	log.Println("DB connected successfully")

	postRepo := repositories.NewPostRepository(db)
	postService := services.NewPostService(postRepo)
	postHandler := handlers.NewUserHandler(postService)

	app := &application{
		config: config,
	}

	router := app.mount(
		routes.PostRouter(postHandler),
	)
	log.Fatal(app.run(router))

}
