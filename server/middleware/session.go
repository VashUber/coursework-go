package middleware

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

var SessionStorage *session.Store

func InitSessionStorage() {
	SessionStorage = session.New(session.Config{
		CookieHTTPOnly: true,
		CookieSecure:   true,
		Expiration:     time.Hour * 36,
	})
}

func SessionChecker(c *fiber.Ctx) error {
	sess, err := SessionStorage.Get(c)
	if err != nil {
		panic(err)
	}

	id := sess.Get("id")

	if id == nil {
		sess.Destroy()
		return c.Redirect("/signin", fiber.StatusUnauthorized)
	}

	return c.Next()
}
