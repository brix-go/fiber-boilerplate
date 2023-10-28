package service

import (
	"github.com/brix-go/fiber/internal/domain/user"
	"github.com/brix-go/fiber/internal/domain/user/dto/requests"
	"github.com/brix-go/fiber/internal/domain/user/dto/responses"
	"github.com/redis/go-redis/v9"
)

type userService struct {
	repo user.UserRepository

	redisClient *redis.Client
}

func NewService(repo user.UserRepository, redis *redis.Client) user.UserService {
	return &userService{
		repo:        repo,
		redisClient: redis,
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
