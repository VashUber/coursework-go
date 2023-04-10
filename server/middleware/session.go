package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

var SessionStorage *session.Store

func InitSessionStorage() {
	SessionStorage = session.New(session.Config{
		CookieHTTPOnly: true,
	})
}

func SessionChecker(c *fiber.Ctx) error {
	sess, err := SessionStorage.Get(c)

	if err != nil {
		panic(err)
	}

	_, ok := sess.Get("id").(string)

	if !ok {
		return c.SendStatus(fiber.StatusForbidden)
	}

	return nil
}
