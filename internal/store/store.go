package store

import (
	"time"

	"os"

	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/storage/postgres/v3"
	"github.com/jackc/pgx/v5/pgxpool"
)

var Store *session.Store
var Environment string

func InitializeStore(dbPool *pgxpool.Pool) {
	// Initialize custom config
	var storage = postgres.New(postgres.Config{
		DB:         dbPool,
		Table:      "user_sessions",
		Reset:      false,
		GCInterval: 10 * time.Second,
	})

	// Session store
	Store = session.New(session.Config{
		CookieHTTPOnly: true,
		Storage:        storage,
		Expiration:     7 * 24 * time.Hour,
	})

	Environment = os.Getenv("ENV")
}
