// Так как я не успеваю всё сделать, а в будущем у меня времени
// на этот проект не будет, то я сделаю бота чисто для себя.

// Если проект не заброшу, то сделаю для всех

package bot

import (
	"log/slog"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/victor-ivanov-ivt20-2/ourdiary/internal/config"
)

var clients = make(map[int64]NotionClientWithSteps)

func Start(log *slog.Logger, token string, env config.OurDiary) error {

	bot, err := tgbotapi.NewBotAPI(token)

	if err != nil {
		return err
	}

	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 60
	updates := bot.GetUpdatesChan(updateConfig)

	for update := range updates {
		chatId := update.Message.Chat.ID
		message := update.Message.Text

		msg, err := Steps(chatId, bot, message, env)

		if err != nil {
			return err
		}

		if err := SendToUser(bot, msg); err != nil {
			return err
		}
	}

	return nil
}

func SendToUser(bot *tgbotapi.BotAPI, msg tgbotapi.MessageConfig) error {
	if _, err := bot.Send(msg); err != nil {
		return err
	}
	return nil
}
