package user

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgxpool"
	queryProvider "github.com/jklq/bug-tracker/internal/db"
	"github.com/jklq/bug-tracker/internal/helpers"
	"github.com/jklq/bug-tracker/internal/store"
	"github.com/jklq/bug-tracker/internal/view"
	"golang.org/x/crypto/bcrypt"
)

// LoginParams holds the structure for login request
type LoginParams struct {
	Email    string `validate:"required,email"`
	Password string `validate:"required"`
}

func handleLoginGet(c *fiber.Ctx, q *queryProvider.Queries, db *pgxpool.Pool) error {
	layout := helpers.HtmxLayoutComponentBasic(c)
	return view.Login(layout, "").Render(c.Context(), c.Response().BodyWriter())
}

// handleLoginPost processes the login request
func handleLoginPost(c *fiber.Ctx, q *queryProvider.Queries, db *pgxpool.Pool) error {
	layout := helpers.HtmxLayoutComponentBasic(c)
	var params LoginParams

	if err := c.BodyParser(&params); err != nil {
		c.Status(fiber.StatusUnauthorized)
		return view.Login(layout, "Invalid credentials").Render(c.Context(), c.Response().BodyWriter())
	}
	user, err := q.GetUserByEmail(c.Context(), params.Email)
	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return view.Login(layout, "Invalid credentials").Render(c.Context(), c.Response().BodyWriter())
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(params.Password)); err != nil {
		c.Status(fiber.StatusUnauthorized)
		return view.Login(layout, "Invalid credentials").Render(c.Context(), c.Response().BodyWriter())
	}

	// Create a new session and save the user ID or other necessary information
	sess, err := store.Store.Get(c)
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return view.Login(layout, "Internal Server Error").Render(c.Context(), c.Response().BodyWriter())
	}

	sess.Set("user_id", user.UserID)

	if err := sess.Save(); err != nil {
		c.Status(fiber.StatusInternalServerError)
		return view.Login(layout, "Internal Server Error").Render(c.Context(), c.Response().BodyWriter())
	}

	return c.Redirect("/app/project")
}