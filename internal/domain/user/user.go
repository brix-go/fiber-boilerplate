package user

import (
	"database/sql"
	"github.com/brix-go/fiber/internal/domain/user/dto/requests"
	"github.com/brix-go/fiber/internal/domain/user/dto/responses"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID         string `gorm:"primary_key"`
	Email      string
	Password   string
	Phone      string
	VerifiedAt sql.NullTime `gorm:"default:null"`
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

func (u *User) FromRegistRequest(req *requests.RegisterRequest) User {
	hashedPass, _ := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.MinCost)
	return User{
		Email:    req.Email,
		Password: string(hashedPass),
		Phone:    req.Phone,
	}
}

func (u *User) ToRegisterResponse() responses.RegisterResponse {
	return responses.RegisterResponse{
		ID:        u.ID,
		Email:     u.Email,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New().String()
	return
}
