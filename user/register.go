package user

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgxpool"
	queryProvider "github.com/jklq/bug-tracker/db"
	"github.com/jklq/bug-tracker/helpers"
	"github.com/jklq/bug-tracker/store"
	"github.com/lucsky/cuid"
	"golang.org/x/crypto/bcrypt"
)

// RegisterParams holds the structure for registration request
type RegisterParams struct {
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required,min=7,max=200"`
	PasswordC string `json:"passwordc"`
	Username  string `json:"username" validate:"required"`
}

// handleRegisterGet renders the registration page
func handleRegisterGet(c *fiber.Ctx, q *queryProvider.Queries, db *pgxpool.Pool) error {
	return c.Render("register", fiber.Map{
		"Title": "Register Page",
	}, "layouts/marketing")
}

// handleRegisterPost processes the registration request
func handleRegisterPost(c *fiber.Ctx, q *queryProvider.Queries, db *pgxpool.Pool) error {
	var params RegisterParams

	if err := c.BodyParser(&params); err != nil {
		return c.Status(fiber.StatusBadRequest).Render("register", fiber.Map{"error": "Invalid request format", "username": params.Username, "email": params.Email}, "layouts/marketing")
	}

	err := helpers.Validate.Struct(params)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).Render("register", fiber.Map{"error": helpers.TranslateError(err, helpers.Translator)[0].Error(), "username": params.Username, "email": params.Email}, "layouts/marketing")
	}

	fmt.Println(params.PasswordC)

	if params.PasswordC != params.Password {
		return c.Status(fiber.StatusBadRequest).Render("register", fiber.Map{"error": "Passwords do not match.", "username": params.Username, "email": params.Email}, "layouts/marketing")
	}

	// Check if the user already exists
	_, err = q.GetUserByEmail(c.Context(), params.Email)
	if err == nil {
		return c.Status(fiber.StatusConflict).Render("register", fiber.Map{"error": "User already exists", "username": params.Username, "email": params.Email}, "layouts/marketing")
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(params.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	// Create new user
	user, err := q.CreateUser(c.Context(), queryProvider.CreateUserParams{
		UserID:       cuid.New(),
		Email:        params.Email,
		Username:     params.Username,
		PasswordHash: string(hashedPassword),
	})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).Render("register", fiber.Map{"error": "Failed to create user", "username": params.Username, "email": params.Email}, "layouts/marketing")
	}

	// Create a new session and save the user ID or other necessary information
	sess, err := store.Store.Get(c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).Render("register", fiber.Map{"error": "Internal Server Error", "username": params.Username, "email": params.Email}, "layouts/marketing")
	}

	sess.Set("user_id", user.UserID)

	if err := sess.Save(); err != nil {
		return c.Status(fiber.StatusInternalServerError).Render("register", fiber.Map{"error": "Internal Server Error", "username": params.Username, "email": params.Email}, "layouts/marketing")
	}

	return c.Redirect("/app")
}
