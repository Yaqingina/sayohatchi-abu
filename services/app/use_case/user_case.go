package use_case

import (
	"go-telegram/services/app/repository"
	"log/slog"
)

type UseCase struct {
	User
}

func New(log *slog.Logger, repo *repository.Repository) *UseCase {
	return &UseCase{
		User: NewUserUseCase(log, repo.User),
	}
}
