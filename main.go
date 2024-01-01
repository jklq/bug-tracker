package main

import (
	"context"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/template/django/v3"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jklq/bug-tracker/dashboard"
	"github.com/jklq/bug-tracker/db"
	"github.com/jklq/bug-tracker/helpers"
	"github.com/jklq/bug-tracker/middleware"
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

	// register functions
	engine.AddFunc("parseDate", helpers.ParseDate)
	engine.AddFunc("parseTime", helpers.ParseTime)
	engine.AddFunc("statusToText", helpers.StatusToText)
	engine.AddFunc("priorityToText", helpers.PriorityToText)

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

	app.Use(logger.New())

	app.Use(middleware.RedirectTrailingSlash)

	app.Get("/", func(c *fiber.Ctx) error {
		// Render index within layouts/main
		return c.Render("index", fiber.Map{
			"title": "Hello, World!",
		}, "layouts/marketing")
	})

	app.Static("/static", "./public")

	userRouter := app.Group("/user")
	dashboardRouter := app.Group("/app")

	user.InitModule(userRouter, queries, dbpool)
	dashboard.InitModule(dashboardRouter, queries, dbpool)

	log.Fatal(app.Listen(":3001"))
}
