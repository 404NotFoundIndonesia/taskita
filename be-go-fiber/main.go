package main

import (
	"database/sql"
	"log"
	"os"
	"taskita/user"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	db, err := sql.Open("mysql", os.Getenv("MYSQL_URI"))
	if err != nil {
		log.Fatalf("database connection error: %v", err)
	}

	// Set up important parts as was told by the documentation.
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	app := fiber.New()
	app.Use(cors.New())

	api := app.Group("/api", logger.New())

	user.New(db, api)

	app.Listen(":8000")
}
