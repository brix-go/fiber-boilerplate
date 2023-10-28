package auth

import (
	"fmt"
	"github.com/brix-go/fiber/internal/domain/user"
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"log"
)

func NewJWThMiddleware(secret string) fiber.Handler {
	log.Println("Secret : ", secret)
	return jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte(secret)},
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			resp := shared.ErrorResponse(contract.ErrUnauthorized, "Unauthorized", err.Error())
			return ctx.Status(401).JSON(resp)
		},
	})
}

func GetCredential(ctx *fiber.Ctx) (err error) {
	defer func() {
		if err := recover(); err != nil {
			ctx.Status(fiber.StatusInternalServerError).JSON(err)
		}
	}()
	useData := ctx.Locals("user").(*jwt.Token)
	claims := useData.Claims.(jwt.MapClaims)
	fmt.Println("CREDENTIALS: ", claims)

	//TODO: Mapping of claims JWT to user's model and save to context local
	credentials := user.User{
		ID:    claims["id"].(string),
		Email: claims["email"].(string),
	}
	ctx.Locals("credentials", credentials)

	return ctx.Next()
}
