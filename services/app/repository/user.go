package repository

import (
	"context"
	"errors"
	"go-telegram/services/app/models"

	"gorm.io/gorm"
)

type User interface {
	Create(ctx context.Context, req models.User) error
	GetByID(ctx context.Context, id int64) (*models.User, error)
	UpdateState(ctx context.Context, userID int64, state uint8) error
	UpdateLanguage(ctx context.Context, userID int64, language string) error
}

type userRepo struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) User {
	return &userRepo{db: db}
}

func (r *userRepo) Create(ctx context.Context, user models.User) error {
	return r.db.WithContext(ctx).Create(&user).Error
}

func (r *userRepo) GetByID(ctx context.Context, id int64) (*models.User, error) {
	var user models.User
	err := r.db.WithContext(ctx).First(&user, "id = ?", id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepo) UpdateState(ctx context.Context, userID int64, state uint8) error {
	result := r.db.WithContext(ctx).Model(&models.User{}).Where("id = ?", userID).Update("state", state)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("not found")
	}
	return nil
}

func (r *userRepo) UpdateLanguage(ctx context.Context, userID int64, language string) error {
	result := r.db.WithContext(ctx).Model(&models.User{}).Where("id = ?", userID).Update("language", language)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("not found")
	}

	return nil
}
