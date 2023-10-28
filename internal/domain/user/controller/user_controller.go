package controller

import (
	"errors"
	"fmt"
	"github.com/brix-go/fiber/internal/domain/user"
	"github.com/brix-go/fiber/internal/domain/user/dto/requests"
	middleware "github.com/brix-go/fiber/middleware/error"
	validation "github.com/brix-go/fiber/middleware/validate"
	"github.com/brix-go/fiber/shared"
	"github.com/gofiber/fiber/v2"
)

type userController struct {
	service user.UserService
}

func NewController(service user.UserService) user.UserController {
	return &userController{
		service: service,
	}
}

func (c *userController) Login(ctx *fiber.Ctx) error {
	var loginReq requests.LoginRequest
	err := ctx.BodyParser(&loginReq)
	if err != nil {
		fmt.Println("ERROR PARSING : ", err)
		return errors.New(shared.ErrInvalidRequestFamily)
	}
	fieldErr := validation.ValidateRequest(loginReq)
	if fieldErr != nil {
		fmt.Println("ERROR : ", fieldErr)
		return errors.New(fieldErr.Error())
	}
	res, err := c.service.Login(&loginReq)
	if err != nil {
		return errors.New(shared.ErrUnexpectedError)
	}
	return middleware.ResponseSuccess(ctx, shared.RespSuccess, res)
}

func (c *userController) Register(ctx *fiber.Ctx) error {
	return middleware.ResponseSuccess(ctx, shared.RespSuccess, nil)

}

func (c *userController) GetDetailUserJWT(ctx *fiber.Ctx) error {
	return middleware.ResponseSuccess(ctx, shared.RespSuccess, nil)

}

func (c *userController) VerifyUser(ctx *fiber.Ctx) error {
	return middleware.ResponseSuccess(ctx, shared.RespSuccess, nil)

}

func (c *userController) ResendOTP(ctx *fiber.Ctx) error {
	return middleware.ResponseSuccess(ctx, shared.RespSuccess, nil)

}
