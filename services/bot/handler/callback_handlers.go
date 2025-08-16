package handler

import (
	"context"
	"log/slog"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (h *Handler) handleCallbacks(ctx context.Context, update tgbotapi.Update) {
	h.log.Info("Handling callback", slog.Int64("user_id", update.CallbackQuery.From.ID))
}
