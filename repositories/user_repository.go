package repositories

import (
	"alexnerotd/apps/crud_user/models"

	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (r *UserRepository) GetUserById(id int) (*models.User, error) {
	var user models.User
	result := r.DB.Take(&user, id)
	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}
