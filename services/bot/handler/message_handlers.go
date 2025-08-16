package handler

import (
	"context"
	"go-telegram/services/bot/store"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (h *Handler) handleMessages(ctx context.Context, update tgbotapi.Update) {
	switch update.Message.Text {
	case store.ButtonLabels[store.LanguageUzLabel][store.UZ], store.ButtonLabels[store.LanguageRuLabel][store.RU], store.ButtonLabels[store.LanguageEnLabel][store.EN]:
		h.handleLanguage(ctx, update)
	default:
		_, _ = h.bot.Send(tgbotapi.NewMessage(update.Message.From.ID, update.Message.Text))
	}
}

func (h *Handler) handleLanguage(ctx context.Context, update tgbotapi.Update) {
	var (
		lang   string
		text   = update.Message.Text
		userID = update.Message.From.ID
	)

	switch text {
	case store.ButtonLabels[store.LanguageUzLabel][store.UZ]:
		lang = store.UZ
	case store.ButtonLabels[store.LanguageRuLabel][store.RU]:
		lang = store.RU
	case store.ButtonLabels[store.LanguageEnLabel][store.EN]:
		lang = store.EN
	default:
		lang = store.UZ
	}

	_, err := h.uc.GetByID(ctx, userID)
	if err != nil {
		text = "Serverda xatolik yuz berdi"
		_, _ = h.bot.Send(tgbotapi.NewMessage(userID, text))
		return
	}

	err = h.uc.User.UpdateLanguage(ctx, userID, lang)
	if err != nil {
		text = "Serverda xatolik yuz berdi"
		_, _ = h.bot.Send(tgbotapi.NewMessage(userID, text))
		return
	}

	err = h.uc.User.UpdateState(ctx, userID, 1)
	if err != nil {
		text = "Serverda xatolik yuz berdi"
		_, _ = h.bot.Send(tgbotapi.NewMessage(userID, text))
		return
	}

	btns := h.keyboard.DashboardKeyboard(lang)
	btns.ResizeKeyboard = true

	msg := tgbotapi.NewMessage(userID, "test")
	msg.ReplyMarkup = btns
	_, _ = h.bot.Send(msg)
	return
}
