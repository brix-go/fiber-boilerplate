package user

import (
	"database/sql"
	"github.com/brix-go/fiber/internal/domain/user/dto/requests"
	"github.com/brix-go/fiber/internal/domain/user/dto/responses"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID         string `gorm:"primary_key"`
	Email      string
	Password   string
	Phone      string       `gorm:"uniqueIndex"`
	VerifiedAt sql.NullTime `gorm:"default:null"`
}

func NewUser(request requests.RegisterRequest) User {
	return User{
		ID:       uuid.NewString(),
		Email:    request.Email,
		Password: request.Password,
		Phone:    request.Phone,
	}
}

type UserRepository interface {
	FindUserByEmail(email string) (*User, error)
	CreateUser(user *User) (*User, error)
	UpdateUser(updatedUser *User) (*User, error)
	GetUserByID(id string) (*User, error)
	GetAllUser() ([]*User, error)
	DeleteUser(id string) error
	VerifyUser(id string) error
}

type UserService interface {
	Login(userRequest *requests.LoginRequest) (res *responses.LoginResponse, err error)
	RegisterUser(userRequest *requests.RegisterRequest) (user *User, err error)
	GetDetailUserById(id string) (user responses.UserDetail, err error)
	VerifyUser(verReq requests.VerifiedUserRequest) error
	ResendOTP(userId string) error
}

type UserController interface {
	Login(ctx *fiber.Ctx) error
	Register(ctx *fiber.Ctx) error
	GetDetailUserJWT(ctx *fiber.Ctx) error
	VerifyUser(ctx *fiber.Ctx) error
	ResendOTP(ctx *fiber.Ctx) error
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New().String()
	return
}
