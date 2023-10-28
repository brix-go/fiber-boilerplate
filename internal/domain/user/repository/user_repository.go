package repository

import (
	"github.com/brix-go/fiber/internal/domain/user"
	"gorm.io/gorm"
	"time"
)

type userRepository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) user.UserRepository {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) FindUserByEmail(email string) (user *user.User, err error) {
	err = r.db.Debug().Take(&user, "email = ?", email).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {

			return nil, nil
		}

		return nil, err
	}

	return user, nil
}

func (r *userRepository) CreateUser(user *user.User) (*user.User, error) {
	err := r.db.Debug().Create(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *userRepository) UpdateUser(updatedUser *user.User) (*user.User, error) {
	return updatedUser, nil
}

func (r *userRepository) GetUserByID(id string) (user *user.User, err error) {
	err = r.db.Debug().First(&user, "id = ?", id).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return user, nil
		}
		return nil, err
	}
	return user, nil
}

func (r *userRepository) GetAllUser() (user []*user.User, err error) {
	return user, nil
}

func (r *userRepository) DeleteUser(id string) error {
	return nil
}

func (r *userRepository) VerifyUser(id string) error {
	err := r.db.Debug().Model(&user.User{}).Where("id = ?", id).Update("verified_at", time.Now()).Error
	if err != nil {
		return err
	}
	return nil
}
