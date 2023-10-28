package auth

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
)

func NewBasicAuthMiddleware() basicauth.Config {
	return basicauth.Config{
		Users: map[string]string{
			"username": "Gojo",
			"password": "Satoru",
		},
		Realm: "Forbidden",
		Authorizer: func(username, password string) bool {
			if username == "Gojo" && password == "Satoru" {
				return true
			}
			return false
		},
		Unauthorized: func(ctx *fiber.Ctx) error {
			resp := dto.ErrorResponse(fiber.StatusUnauthorized, "Forbidden, unauthorized")
			return ctx.Status(fiber.StatusUnauthorized).JSON(resp)
		},
		ContextPassword: "_pass",
		ContextUsername: "_user",
	}
}
