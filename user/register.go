package user

import (
	"fmt"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgxpool"
	queryProvider "github.com/jklq/bug-tracker/db"
	"github.com/jklq/bug-tracker/store"
	"golang.org/x/crypto/bcrypt"
)

// RegisterParams holds the structure for registration request
type RegisterParams struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=7,max=200"`
	Username string `json:"username" validate:"required"`
}

var validate = validator.New()

var english = en.New()
var uni = ut.New(english, english)
var trans, _ = uni.GetTranslator("en")

var _ = enTranslations.RegisterDefaultTranslations(validate, trans)

func translateError(err error, trans ut.Translator) (errs []error) {
	if err == nil {
		return nil
	}
	validatorErrs := err.(validator.ValidationErrors)
	for _, e := range validatorErrs {
		translatedErr := fmt.Errorf(e.Translate(trans))
		errs = append(errs, translatedErr)
	}
	return errs
}

// handleRegisterGet renders the registration page
func handleRegisterGet(c *fiber.Ctx, q *queryProvider.Queries, db *pgxpool.Pool) error {
	return c.Render("register", fiber.Map{
		"Title": "Register Page",
	})
}

// handleRegisterPost processes the registration request
func handleRegisterPost(c *fiber.Ctx, q *queryProvider.Queries, db *pgxpool.Pool) error {
	var params RegisterParams

	if err := c.BodyParser(&params); err != nil {
		return c.Status(fiber.StatusBadRequest).Render("register", fiber.Map{"error": "Invalid request format"})
	}

	err := validate.Struct(params)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).Render("register", fiber.Map{"error": translateError(err, trans)[0].Error()})
	}

	// Check if the user already exists
	_, err = q.GetUserByEmail(c.Context(), params.Email)
	if err == nil {
		return c.Status(fiber.StatusConflict).Render("register", fiber.Map{"error": "User already exists"})
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(params.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	// Create new user
	user, err := q.CreateUser(c.Context(), queryProvider.CreateUserParams{
		Email:    params.Email,
		Username: params.Username,

		PasswordHash: string(hashedPassword),
	})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).Render("register", fiber.Map{"error": "Failed to create user"})
	}

	// Create a new session and save the user ID or other necessary information
	sess, err := store.Store.Get(c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).Render("register", fiber.Map{"error": "Internal Server Error"})
	}

	sess.Set("user_id", user.UserID)

	if err := sess.Save(); err != nil {
		return c.Status(fiber.StatusInternalServerError).Render("register", fiber.Map{"error": "Internal Server Error"})
	}

	return c.Redirect("/")
}
