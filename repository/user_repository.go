package repository

import (
	"travelnest/model"

	"gorm.io/gorm"
)

type UserRepository struct {
    connection *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
    return UserRepository{
		connection: db,
	}
}

func (r *UserRepository) GetByUsername(username string) (*model.User, error) {
    var user model.User
    if err := r.connection.Where("username = ?", username).First(&user).Error; err != nil {
        return nil, err
    }
    return &user, nil
}
