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
		{store.ButtonLabels[store.LabelLanguageUz][store.UZ]},
		{store.ButtonLabels[store.LabelLanguageRu][store.RU]},
		{store.ButtonLabels[store.LabelLanguageEn][store.EN]},
	}

	return k.generateReplyKeyboard(buttons)
}

func (k *Keyboard) DashboardKeyboard(lang string) tgbotapi.ReplyKeyboardMarkup {
	buttons := [][]string{
		{store.ButtonLabels[store.LabelAdvertisement][lang], store.ButtonLabels[store.LabelAdvice][lang]},
		{store.ButtonLabels[store.LabelFaq][lang]},
		{store.ButtonLabels[store.LabelSuggestion][lang]},
		{store.ButtonLabels[store.LabelYoutube][lang], store.ButtonLabels[store.LabelTelegram][lang]},
		{store.ButtonLabels[store.LabelInstagram][lang], store.ButtonLabels[store.LabelFacebook][lang]},
		{store.ButtonLabels[store.LabelReturn][lang]},
	}

	return k.generateReplyKeyboard(buttons)
}

func (k *Keyboard) AdvertisementKeyboard(lang string) tgbotapi.ReplyKeyboardMarkup {
	buttons := [][]string{
		{store.ButtonLabels[store.LabelPrice][lang], store.ButtonLabels[store.LabelStatistics][lang]},
		{store.ButtonLabels[store.LabelReturn][lang]},
	}

	return k.generateReplyKeyboard(buttons)
}

func (k *Keyboard) AdviceKeyboard(lang string) tgbotapi.ReplyKeyboardMarkup {
	buttons := [][]string{
		{store.ButtonLabels[store.LabelFree][lang], store.ButtonLabels[store.LabelPaid][lang]},
		{store.ButtonLabels[store.LabelReturn][lang]},
	}

	return k.generateReplyKeyboard(buttons)
}

func (k *Keyboard) SuggestionKeyboard(lang string) tgbotapi.ReplyKeyboardMarkup {
	buttons := [][]string{
		{store.ButtonLabels[store.LabelPhone][lang], store.ButtonLabels[store.LabelTelegram][lang]},
		{store.ButtonLabels[store.LabelReturn][lang]},
	}

	return k.generateReplyKeyboard(buttons)
}
