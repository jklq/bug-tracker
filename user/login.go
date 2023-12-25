package user

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgxpool"
	queryProvider "github.com/jklq/bug-tracker/db"
	"github.com/jklq/bug-tracker/store"
	"golang.org/x/crypto/bcrypt"
)

// LoginParams holds the structure for login request
type LoginParams struct {
	Email    string `validate:"required,email"`
	Password string `validate:"required"`
}

// handleLoginGet renders the login page
func handleLoginGet(c *fiber.Ctx, q *queryProvider.Queries, db *pgxpool.Pool) error {
	return c.Render("login", fiber.Map{
		"Title": "Login Page",
	}, "layouts/marketing")
}

// handleLoginPost processes the login request
func handleLoginPost(c *fiber.Ctx, q *queryProvider.Queries, db *pgxpool.Pool) error {
	var params LoginParams

	if err := c.BodyParser(&params); err != nil {
		return c.Status(fiber.StatusBadRequest).Render("login", fiber.Map{"error": "Invalid request format"})
	}
	user, err := q.GetUserByEmail(c.Context(), params.Email)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).Render("login", fiber.Map{"error": "Invalid credentials"})
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(params.Password)); err != nil {
		return c.Status(fiber.StatusUnauthorized).Render("login", fiber.Map{"error": "Invalid credentials"})
	}

	// Create a new session and save the user ID or other necessary information
	sess, err := store.Store.Get(c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).Render("login", fiber.Map{"error": "Internal Server Error"})
	}

	sess.Set("user_id", user.UserID)

	if err := sess.Save(); err != nil {
		return c.Status(fiber.StatusInternalServerError).Render("login", fiber.Map{"error": "Internal Server Error"})
	}

	return c.Redirect("/app")
}
