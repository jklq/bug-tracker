package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"

	"github.com/gofiber/template/html/v2"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jklq/project-tracker/internal/db"
	"github.com/jklq/project-tracker/internal/helpers"
	"github.com/jklq/project-tracker/internal/middleware"
	"github.com/jklq/project-tracker/internal/project"
	"github.com/jklq/project-tracker/internal/store"
	"github.com/jklq/project-tracker/internal/ticket"
	"github.com/jklq/project-tracker/internal/user"
	"github.com/jklq/project-tracker/internal/view"
	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load(".env")
}
func main() {
	engine := html.New("./internal/view", ".html")

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

	app.Static("/static", "./public")

	app.Use(middleware.EnsureHtmlContentType)

	app.Get("/", func(c *fiber.Ctx) error {
		layout := helpers.HtmxLayoutComponentBasic(c)
		return view.Index(layout).Render(c.Context(), c.Response().BodyWriter())
	})

	app.Get("/app", func(c *fiber.Ctx) error { return c.Redirect("/app/project") })

	userRouter := app.Group("/user")
	projectRouter := app.Group("/app/project")
	ticketRouter := app.Group("/app")

	user.InitModule(userRouter, queries, dbpool)
	project.InitModule(projectRouter, queries, dbpool)
	ticket.InitModule(ticketRouter, queries, dbpool)

	log.Fatal(app.Listen(fmt.Sprintf(":%s", os.Getenv("PORT"))))
}
