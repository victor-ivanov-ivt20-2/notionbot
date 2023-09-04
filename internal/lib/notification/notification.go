package notification

import (
	"github.com/go-co-op/gocron"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/victor-ivanov-ivt20-2/ourdiary/internal/bot"
	"github.com/victor-ivanov-ivt20-2/ourdiary/internal/notion"
)

func NotificateMe(tgbot *tgbotapi.BotAPI, chatId int64, client notion.NotionClient, scheduler *gocron.Scheduler, title string, lessonStartTime string, room string) error {
	_, err := scheduler.Every(15).Seconds().Do(func(title string, lessonStartTime string, room string) error {
		schedulerMessage := tgbotapi.NewMessage(chatId, title+" начнётся в "+lessonStartTime+" в "+room)
		if err := bot.SendToUser(tgbot, schedulerMessage); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}
