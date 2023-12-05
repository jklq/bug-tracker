package main

import (
	"context"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/django/v3"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jklq/bug-tracker/db"
	"github.com/jklq/bug-tracker/store"
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

	store.InitializeStore(dbpool)

	queries := db.New(dbpool)

	defer dbpool.Close()

	app := fiber.New(fiber.Config{
		Views:       engine,
		ViewsLayout: "layouts/main",
	})

	app.Get("/", func(c *fiber.Ctx) error {
		// Render index within layouts/main
		return c.Render("index", fiber.Map{
			"title": "Hello, World!",
		})
	})

	app.Static("/static", "./public")

	radarRouter := app.Group("/user")

	user.InitModule(radarRouter, queries, dbpool)

	log.Fatal(app.Listen(":3001"))
}
