package main

import (
	"context"
	"go-telegram/config"
	"go-telegram/pkgs/database"
	"go-telegram/services/app/repository"
	"go-telegram/services/app/use_case"
	botHandler "go-telegram/services/bot/handler"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {

	logger := slog.New(slog.NewJSONHandler(
		os.Stdout,
		&slog.HandlerOptions{AddSource: false}),
	)

	repo := repository.New(database.Init())
	uc := use_case.New(logger, repo)

	botAPI, err := tgbotapi.NewBotAPI(config.Config("APP_BOT_TOKEN"))
	if err != nil {
		logger.Error("Failed to create new bot", slog.String("error", err.Error()))
		os.Exit(1)
	}

	bot := botHandler.New(logger, uc, botAPI)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	bot.Run(ctx)

	app := fiber.New()
	app.Use(cors.New(cors.Config{AllowOrigins: "*", AllowHeaders: "*"}))
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.SendString("OK")
	})

	go func() {
		port := config.Config("APP_PORT")
		logger.Info("HTTP server listening", slog.String("port", port))

		if err := app.Listen(":" + port); err != nil {
			logger.Error("Failed to start HTTP server", slog.String("error", err.Error()))
			cancel()
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	logger.Info("Shutdown signal received")
	cancel()

	if err := app.Shutdown(); err != nil {
		logger.Error("Fiber server forced to shutdown", slog.String("error", err.Error()))
	}

	logger.Info("Application stopped gracefully")
}
