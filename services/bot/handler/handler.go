package handler

import (
	"context"
	"go-telegram/services/app/use_case"
	"go-telegram/services/bot/keyboard"
	"go-telegram/services/bot/store"
	"log/slog"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Handler struct {
	log      *slog.Logger
	uc       *use_case.UseCase
	bot      *tgbotapi.BotAPI
	keyboard *keyboard.Keyboard
	store    store.Store
}

func New(logger *slog.Logger, useCase *use_case.UseCase, bot *tgbotapi.BotAPI) Handler {
	return Handler{
		log:      logger,
		uc:       useCase,
		bot:      bot,
		keyboard: keyboard.New(),
		store:    store.New(),
	}
}

func (h *Handler) Run(ctx context.Context) {
	updatesChan := make(chan tgbotapi.Update, 100)

	numWorkers := 10
	for i := 0; i < numWorkers; i++ {
		go func(workerID int) {
			for upd := range updatesChan {
				h.HandleUpdate(ctx, upd)
			}
		}(i + 1)
	}

	go func() {
		u := tgbotapi.NewUpdate(0)
		u.Timeout = 60

		h.log.Info("Bot started")

		for {
			select {
			case <-ctx.Done():
				h.log.Info("Bot shutdown requested")
				close(updatesChan)
				return
			default:
				updates, err := h.bot.GetUpdates(u)
				if err != nil {
					h.log.Error("Failed to fetch updates", slog.String("error", err.Error()))
					time.Sleep(1 * time.Second)
					continue
				}

				for _, upd := range updates {
					if upd.UpdateID >= u.Offset {
						u.Offset = upd.UpdateID + 1
						updatesChan <- upd
					}
				}
			}
		}
	}()
}

func (h *Handler) HandleUpdate(ctx context.Context, update tgbotapi.Update) {
	if update.CallbackQuery != nil {
		h.handleCallbacks(ctx, update)
	} else if update.Message != nil && update.Message.IsCommand() {
		h.handleCommands(ctx, update)
	} else if update.Message != nil {
		h.handleMessages(ctx, update)
	}
}
