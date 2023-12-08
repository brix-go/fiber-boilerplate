package service

import (
	"github.com/brix-go/fiber/infrastructure/redis"
	"github.com/brix-go/fiber/internal/domain/user"
	"github.com/brix-go/fiber/internal/domain/user/dto/requests"
	"github.com/brix-go/fiber/internal/domain/user/dto/responses"
)

type userService struct {
	repo user.UserRepository

	redisRepo *redis_client.Repository
}

func NewService(repo user.UserRepository, redis *redis_client.Repository) user.UserService {
	return &userService{
		repo:      repo,
		redisRepo: redis,
	}
}

func (s *userService) Login(userRequest *requests.LoginRequest) (*responses.LoginResponse, error) {
	//TODO: Implement login
	return &responses.LoginResponse{}, nil
}

func (s *userService) RegisterUser(userRequest *requests.RegisterRequest) (userData *user.User, err error) {
	//TODO: Implement register
	return userData, nil
}

func (s *userService) GetDetailUserById(id string) (res responses.UserDetail, err error) {
	//TODO: Implement getDetailUserbyId
	return res, nil
}

func (s *userService) VerifyUser(verReq requests.VerifiedUserRequest) error {
	// TODO: Implement verify user
	return nil
}

func (s *userService) ResendOTP(userId string) error {
	return nil
}
