package repository

import (
	"backend/app/models"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func (u *UserRepository) Create(user models.User) error {
	u.db.Create(&user)
	return nil
}
