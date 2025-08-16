package keyboard

import (
	"go-telegram/services/bot/store"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Keyboard struct{}

func New() *Keyboard {
	return &Keyboard{}
}

func (k *Keyboard) generateReplyKeyboard(buttons [][]string) tgbotapi.ReplyKeyboardMarkup {
	var keyboard [][]tgbotapi.KeyboardButton
	for _, row := range buttons {
		var kbRow []tgbotapi.KeyboardButton
		for _, text := range row {
			kbRow = append(kbRow, tgbotapi.NewKeyboardButton(text))
		}
		keyboard = append(keyboard, kbRow)
	}

	return tgbotapi.NewReplyKeyboard(keyboard...)
}

func (k *Keyboard) LanguageKeyboard() tgbotapi.ReplyKeyboardMarkup {
	buttons := [][]string{
		{store.ButtonLabels[store.LanguageUzLabel][store.UZ]},
		{store.ButtonLabels[store.LanguageRuLabel][store.RU]},
		{store.ButtonLabels[store.LanguageEnLabel][store.EN]},
	}

	return k.generateReplyKeyboard(buttons)
}

func (k *Keyboard) DashboardKeyboard(lang string) tgbotapi.ReplyKeyboardMarkup {
	buttons := [][]string{
		{store.ButtonLabels[store.AdvertisementLabel][lang], store.ButtonLabels[store.AdviceLabel][lang]},
		{store.ButtonLabels[store.FaqLabel][lang]},
		{store.ButtonLabels[store.SuggestionLabel][lang]},
		{store.ButtonLabels[store.YoutubeLabel][lang], store.ButtonLabels[store.TelegramLabel][lang]},
		{store.ButtonLabels[store.InstagramLabel][lang], store.ButtonLabels[store.FacebookLabel][lang]},
		{store.ButtonLabels[store.ReturnLabel][lang]},
	}

	return k.generateReplyKeyboard(buttons)
}
