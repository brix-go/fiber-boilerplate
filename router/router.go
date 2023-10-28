package router

import (
	"github.com/brix-go/fiber/internal/domain/user"
	"github.com/gofiber/fiber/v2"
)

type RouteParams struct {
	user.UserController
}

type RouterStruct struct {
	RouteParams RouteParams
}

func NewRouter(params *RouteParams) RouterStruct {
	return RouterStruct{
		RouteParams: *params,
	}
}

func (r *RouterStruct) SetupRoute(app *fiber.App) {
	v1 := app.Group("/api/v1")
	v1.Route("/auth", func(router fiber.Router) {
		router.Post("/login", r.RouteParams.UserController.Login)
	})
}
