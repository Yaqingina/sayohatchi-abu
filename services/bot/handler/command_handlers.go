package handler

import (
	"context"
	"go-telegram/services/app/models"
	"go-telegram/services/bot/store"
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
	userID := update.Message.Chat.ID

	text := "Tilni tanlang\nВыберите язык\nChoose language"
	user, err := h.uc.User.GetByID(ctx, userID)
	if err != nil {
		h.log.Error("failed to fetch user", slog.Int64("user_id", userID), slog.Any("error", err))
		_, _ = h.bot.Send(tgbotapi.NewMessage(userID, err.Error()))
		return
	}

	if user != nil {
		err = h.uc.User.UpdateState(ctx, userID, store.StateLanguage)
		if err != nil {
			text = "Serverda xatolik yuz berdi"
			_, _ = h.bot.Send(tgbotapi.NewMessage(userID, text))
			return
		}

		text = "Botga xush kelibsiz"
		msg := tgbotapi.NewMessage(userID, text)

		btns := h.keyboard.LanguageKeyboard()
		btns.ResizeKeyboard = true
		btns.OneTimeKeyboard = true

		msg.ReplyMarkup = btns
		_, _ = h.bot.Send(msg)
		return
	}

	newUser := models.User{
		ID:        userID,
		FirstName: update.Message.From.FirstName,
		LastName:  update.Message.From.LastName,
		UserName:  update.Message.From.LastName,
		State:     store.StateLanguage,
	}

	if err := h.uc.Create(ctx, newUser); err != nil {
		log.Error("failed to create user", slog.Int64("user_id", userID), slog.Any("error", err))
		_, _ = h.bot.Send(tgbotapi.NewMessage(userID, err.Error()))
		return
	}

	msg := tgbotapi.NewMessage(userID, text)

	btns := h.keyboard.LanguageKeyboard()
	btns.ResizeKeyboard = true
	btns.OneTimeKeyboard = true

	msg.ReplyMarkup = btns
	_, _ = h.bot.Send(msg)
	return
}
