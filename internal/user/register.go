package user

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgxpool"
	queryProvider "github.com/jklq/project-tracker/internal/db"
	"github.com/jklq/project-tracker/internal/helpers"
	"github.com/jklq/project-tracker/internal/store"
	"github.com/jklq/project-tracker/internal/view"
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
	layout := helpers.HtmxLayoutComponentBasic(c)

	return view.Register(layout, "", view.RegisterParams{}).Render(c.Context(), c.Response().BodyWriter())
}

// handleRegisterPost processes the registration request
func handleRegisterPost(c *fiber.Ctx, q *queryProvider.Queries, db *pgxpool.Pool) error {
	layout := helpers.HtmxLayoutComponentBasic(c)
	var params RegisterParams

	if err := c.BodyParser(&params); err != nil {
		c.Status(fiber.StatusBadRequest)
		return view.Register(layout, "Invalid request format", view.RegisterParams{}).Render(c.Context(), c.Response().BodyWriter())
	}

	viewParams := view.RegisterParams{Email: params.Email, Username: params.Username}

	err := helpers.Validate.Struct(params)
	if err != nil {
		errorMsg := helpers.TranslateError(err, helpers.Translator)[0].Error()

		c.Status(fiber.StatusBadRequest)
		return view.Register(layout, errorMsg, viewParams).Render(c.Context(), c.Response().BodyWriter())
	}

	if params.PasswordC != params.Password {
		c.Status(fiber.StatusBadRequest)
		return view.Register(layout, "Passwords do not match", viewParams).Render(c.Context(), c.Response().BodyWriter())
	}

	// Check if the user already exists
	_, err = q.GetUserByEmail(c.Context(), params.Email)
	if err == nil {
		c.Status(fiber.StatusConflict)
		return view.Register(layout, "User already exists", viewParams).Render(c.Context(), c.Response().BodyWriter())
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(params.Password), bcrypt.DefaultCost)
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return view.Register(layout, "Failed to create user", viewParams).Render(c.Context(), c.Response().BodyWriter())
	}

	// Create new user
	user, err := q.CreateUser(c.Context(), queryProvider.CreateUserParams{
		UserID:       cuid.New(),
		Email:        params.Email,
		Username:     params.Username,
		PasswordHash: string(hashedPassword),
	})
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return view.Register(layout, "Failed to create user", viewParams).Render(c.Context(), c.Response().BodyWriter())
	}

	// Create a new session and save the user ID or other necessary information
	sess, err := store.Store.Get(c)
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return view.Register(layout, "Failed to create user", viewParams).Render(c.Context(), c.Response().BodyWriter())
	}

	sess.Set("user_id", user.UserID)

	if err := sess.Save(); err != nil {
		c.Status(fiber.StatusInternalServerError)
		return view.Register(layout, "Failed to create user", viewParams).Render(c.Context(), c.Response().BodyWriter())
	}

	return c.Redirect("/app")
}
