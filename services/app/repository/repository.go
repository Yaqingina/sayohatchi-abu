package repository

import (
	"gorm.io/gorm"
)

type Repository struct {
	User
}

func New(db *gorm.DB) *Repository {
	return &Repository{
		User: NewUserRepository(db),
	}
}
