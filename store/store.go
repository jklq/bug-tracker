package store

import "github.com/gofiber/fiber/v2/middleware/session"

// Session store
var Store = session.New(session.Config{
	CookieHTTPOnly: true,
})
