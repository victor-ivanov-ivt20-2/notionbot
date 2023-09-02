package main

import (
	"fmt"
	"log/slog"
	"os"
	"time"

	"github.com/victor-ivanov-ivt20-2/ourdiary/internal/bot"
	"github.com/victor-ivanov-ivt20-2/ourdiary/internal/config"
	"github.com/victor-ivanov-ivt20-2/ourdiary/internal/lib/logger/sl"
	"github.com/victor-ivanov-ivt20-2/ourdiary/internal/storage/sqlite"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func init() {
	loc, err := time.LoadLocation("Asia/Yakutsk")
	if err != nil {
		fmt.Println(err)
		return
	}
	time.Local = loc
}

func main() {
	cfg := config.MustLoad()

	log := setupLogger(cfg.Env)

	log.Info("starting ourdiary", slog.String("env", cfg.Env))
	log.Debug("debug messages are enabled")

	storage, err := sqlite.New(cfg.StoragePath)
	if err != nil {
		log.Error("failed to init storage", sl.Err(err))
		os.Exit(1)
	}

	_ = storage

	if err := bot.Start(log, cfg.TelegramBotToken, cfg.OurDiary); err != nil {
		log.Error("tgbot error", sl.Err(err))
	}
}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger
	switch env {
	case envLocal:
		log = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case envDev:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case envProd:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	}
	return log
}
