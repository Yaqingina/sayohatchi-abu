package handler

import (
	"context"
	"go-telegram/services/bot/store"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (h *Handler) handleMessages(ctx context.Context, update tgbotapi.Update) {
	switch update.Message.Text {
	case store.ButtonLabels[store.LabelLanguageUz][store.UZ], store.ButtonLabels[store.LabelLanguageRu][store.RU], store.ButtonLabels[store.LabelLanguageEn][store.EN]:
		h.handleLanguage(ctx, update)
	case store.ButtonLabels[store.LabelAdvertisement][store.UZ], store.ButtonLabels[store.LabelAdvertisement][store.RU], store.ButtonLabels[store.LabelAdvertisement][store.EN]:
		h.handleAdvertisement(ctx, update)
	case store.ButtonLabels[store.LabelAdvice][store.UZ], store.ButtonLabels[store.LabelAdvice][store.RU], store.ButtonLabels[store.LabelAdvice][store.EN]:
		h.handleAdvertisement(ctx, update)
	case store.ButtonLabels[store.LabelSuggestion][store.UZ], store.ButtonLabels[store.LabelSuggestion][store.RU], store.ButtonLabels[store.LabelSuggestion][store.EN]:
		h.handleAdvertisement(ctx, update)
	default:
		if update.Message.Text == "reset" {
			h.startCommand(ctx, update)
			return
		}
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
	case store.ButtonLabels[store.LabelLanguageUz][store.UZ]:
		lang = store.UZ
	case store.ButtonLabels[store.LabelLanguageRu][store.RU]:
		lang = store.RU
	case store.ButtonLabels[store.LabelLanguageEn][store.EN]:
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

	err = h.uc.User.UpdateState(ctx, userID, store.StateMenu)
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

func (h *Handler) handleAdvertisement(ctx context.Context, update tgbotapi.Update) {
	var (
		text   = update.Message.Text
		userID = update.Message.From.ID
	)

	user, err := h.uc.GetByID(ctx, userID)
	if err != nil {
		text = "Serverda xatolik yuz berdi"
		_, _ = h.bot.Send(tgbotapi.NewMessage(userID, text))
		return
	}

	btns := h.keyboard.AdvertisementKeyboard(user.Language)
	btns.ResizeKeyboard = true

	msg := tgbotapi.NewMessage(userID, "test")
	msg.ReplyMarkup = btns
	_, _ = h.bot.Send(msg)
	return
}

func (h *Handler) handleAdvice(ctx context.Context, update tgbotapi.Update) {
	var (
		text   = update.Message.Text
		userID = update.Message.From.ID
	)

	user, err := h.uc.GetByID(ctx, userID)
	if err != nil {
		text = "Serverda xatolik yuz berdi"
		_, _ = h.bot.Send(tgbotapi.NewMessage(userID, text))
		return
	}

	btns := h.keyboard.AdviceKeyboard(user.Language)
	btns.ResizeKeyboard = true

	msg := tgbotapi.NewMessage(userID, "test")
	msg.ReplyMarkup = btns
	_, _ = h.bot.Send(msg)
	return
}

func (h *Handler) handleSuggestion(ctx context.Context, update tgbotapi.Update) {
	var (
		text   = update.Message.Text
		userID = update.Message.From.ID
	)

	user, err := h.uc.GetByID(ctx, userID)
	if err != nil {
		text = "Serverda xatolik yuz berdi"
		_, _ = h.bot.Send(tgbotapi.NewMessage(userID, text))
		return
	}

	btns := h.keyboard.SuggestionKeyboard(user.Language)
	btns.ResizeKeyboard = true

	msg := tgbotapi.NewMessage(userID, "test")
	msg.ReplyMarkup = btns
	_, _ = h.bot.Send(msg)
	return
}
