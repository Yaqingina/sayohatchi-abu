package use_case

import (
	"context"
	"go-telegram/services/app/models"
	"go-telegram/services/app/repository"
	"log/slog"
)

type User interface {
	Create(ctx context.Context, req models.User) error
	GetByID(ctx context.Context, id int64) (*models.User, error)
	UpdateState(ctx context.Context, userID int64, state uint8) error
	UpdateLanguage(ctx context.Context, userID int64, language string) error
}

type userUseCase struct {
	log *slog.Logger
	r   repository.User
}

func NewUserUseCase(log *slog.Logger, repository repository.User) User {
	return &userUseCase{log: log, r: repository}
}

func (s *userUseCase) Create(ctx context.Context, req models.User) error {
	return s.r.Create(ctx, req)
}

func (s *userUseCase) GetByID(ctx context.Context, id int64) (*models.User, error) {
	return s.r.GetByID(ctx, id)
}

func (s *userUseCase) UpdateState(ctx context.Context, userID int64, state uint8) error {
	return s.r.UpdateState(ctx, userID, state)
}

func (s *userUseCase) UpdateLanguage(ctx context.Context, userID int64, language string) error {
	return s.r.UpdateLanguage(ctx, userID, language)
}
