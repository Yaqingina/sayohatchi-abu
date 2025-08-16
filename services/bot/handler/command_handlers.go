package handler

import (
	"context"
	"go-telegram/services/app/models"
	"log/slog"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/gofiber/fiber/v2/log"
)

func (h *Handler) handleCommands(ctx context.Context, update tgbotapi.Update) {
	switch update.Message.Command() {
	case "start":
		h.startCommand(ctx, update)
	}
}

func (h *Handler) startCommand(ctx context.Context, update tgbotapi.Update) {
	chatID := update.Message.Chat.ID

	text := "Tilni tanlang\nВыберите язык\nChoose language"
	user, err := h.uc.User.GetByID(ctx, chatID)
	if err != nil {
		h.log.Error("failed to fetch user", slog.Int64("user_id", chatID), slog.Any("error", err))
		_, _ = h.bot.Send(tgbotapi.NewMessage(chatID, err.Error()))
		return
	}

	if user != nil {
		_, _ = h.bot.Send(tgbotapi.NewMessage(chatID, "Botga xush kelibsiz"))
		return
	}

	newUser := models.User{
		ID:        chatID,
		FirstName: update.Message.From.FirstName,
		LastName:  update.Message.From.LastName,
		UserName:  update.Message.From.LastName,
	}

	if err := h.uc.Create(context.Background(), newUser); err != nil {
		log.Error("failed to create user", slog.Int64("user_id", chatID), slog.Any("error", err))
		_, _ = h.bot.Send(tgbotapi.NewMessage(chatID, err.Error()))
		return
	}

	msg := tgbotapi.NewMessage(chatID, text)

	btns := h.keyboard.LanguageKeyboard()
	btns.ResizeKeyboard = true
	btns.OneTimeKeyboard = true

	msg.ReplyMarkup = btns
	_, _ = h.bot.Send(msg)
	return
}
