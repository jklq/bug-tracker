package main

import (
	"context"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"

	"github.com/gofiber/template/html/v2"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jklq/bug-tracker/db"
	"github.com/jklq/bug-tracker/helpers"
	"github.com/jklq/bug-tracker/middleware"
	"github.com/jklq/bug-tracker/project"
	"github.com/jklq/bug-tracker/store"
	"github.com/jklq/bug-tracker/ticket"
	"github.com/jklq/bug-tracker/user"
	"github.com/jklq/bug-tracker/view"
	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load(".env")
}
func main() {

	// Initialize standard Go html template engine
	// engine := django.New("./views", ".django")

	engine := html.New("./view", ".html")

	// register functions
	//engine.AddFunc("parseDate", helpers.ParseDate)
	//engine.AddFunc("parseTime", helpers.ParseTime)
	//engine.AddFunc("statusToText", helpers.StatusToText)
	//engine.AddFunc("priorityToText", helpers.PriorityToText)

	//engine.AddFunc("c", helpers.CleanHTML)

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

	log.Fatal(app.Listen(":3001"))
}
