package main

import (
	"context"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/django/v3"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jklq/bug-tracker/db"
	"github.com/jklq/bug-tracker/user"
	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load(".env")
}
func main() {
	// Initialize standard Go html template engine
	engine := django.New("./views", ".django")

	// Init database pool
	dbpool, err := pgxpool.New(context.Background(), os.Getenv("DATABASE_URL"))

	if err != nil {
		panic(err)
	}

	queries := db.New(dbpool)

	defer dbpool.Close()

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Get("/", func(c *fiber.Ctx) error {
		// Render index within layouts/main
		return c.Render("index", fiber.Map{
			"title": "Hello, World!",
		}, "layouts/main")
	})

	app.Static("/", "./public")
	app.Static("/", "./favicon")

	radarRouter := app.Group("/user")

	user.InitRadar(radarRouter, queries)

	log.Fatal(app.Listen(":3001"))
}
