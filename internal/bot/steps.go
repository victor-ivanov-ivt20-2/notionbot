package bot

import (
	"time"

	"github.com/go-co-op/gocron"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/jomei/notionapi"
	"github.com/victor-ivanov-ivt20-2/ourdiary/internal/config"
	"github.com/victor-ivanov-ivt20-2/ourdiary/internal/notion"
)

type Step string

const (
	WELCOME    Step = "welcome"
	LOGIN      Step = "login"
	SET_PERSON Step = "setPerson"
	WORKING    Step = "working"
)

var workingKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Обновить расписание"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Всё расписание"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Расписание на завтра"),
		tgbotapi.NewKeyboardButton("Расписание на сегодня"),
	),
)

func Steps(chatId int64, bot *tgbotapi.BotAPI, scheduler *gocron.Scheduler, message string, env config.OurDiary) (tgbotapi.MessageConfig, error) {
	var msg tgbotapi.MessageConfig
	client, err := GetClient(chatId)

	if err != nil {
		return tgbotapi.MessageConfig{}, err
	}
	if client.CurrentStep != WORKING {
		msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
	}
	switch client.CurrentStep {
	case WELCOME:
		msg = tgbotapi.NewMessage(chatId, "Это пока всё что может этот бот :c")
		SetStep(chatId, client, LOGIN)
	case LOGIN:
		if message == env.Password {
			client.NotionClient.Notion = notion.SetClient(notionapi.Token(env.Token))
			clients[chatId] = client
			msg = tgbotapi.NewMessage(chatId, "Или ...")
			SetStep(chatId, client, SET_PERSON)
		} else {
			msg = tgbotapi.NewMessage(chatId, "Это пока всё что может этот бот :c")
		}
	case SET_PERSON:
		if message != env.First.Email && message != env.Second.Email {
			msg = tgbotapi.NewMessage(chatId, "Повторяй")
		} else {
			client.NotionClient.TasksId = env.TasksId
			client.NotionClient.ScheduleId = env.ScheduleId
			if message == env.First.Email {
				client.NotionClient.Email = env.First.Email
				client.NotionClient.PageId = env.First.PageId
				client.NotionClient.UserId = env.First.UserId
			} else if message == env.Second.Email {
				client.NotionClient.Email = env.Second.Email
				client.NotionClient.PageId = env.Second.PageId
				client.NotionClient.UserId = env.Second.UserId
			}
			msg = tgbotapi.NewMessage(chatId, "Поздравляю!")
			msg.ReplyMarkup = workingKeyboard
			SetStep(chatId, client, WORKING)
		}

	case WORKING:

		// schedulerTask := func(title string, lessonStartTime string, room string) error {
		// 	schedulerMessage := tgbotapi.NewMessage(chatId, title+" начнётся в "+lessonStartTime+" в "+room)
		// 	if err := SendToUser(bot, schedulerMessage); err != nil {
		// 		return err
		// 	}
		// 	return nil
		// }

		switch message {
		case "Обновить расписание":
			if err := notion.UpdateSchedule(client.NotionClient); err != nil {
				return tgbotapi.MessageConfig{}, err
			}
			msg = tgbotapi.NewMessage(chatId, "Расписание на вашей странице было обновлено!")
		case "Всё расписание":
			answer, err := notion.GetAllSchedule(client.NotionClient)
			if err != nil {
				return tgbotapi.MessageConfig{}, err
			}
			msg = tgbotapi.NewMessage(chatId, answer)
		case "Расписание на завтра":
			answer, err := notion.GetScheduleForDay(client.NotionClient, time.Now().AddDate(0, 0, 1))
			if err != nil {
				return tgbotapi.MessageConfig{}, err
			}
			msg = tgbotapi.NewMessage(chatId, answer)
		case "Расписание на сегодня":
			answer, err := notion.GetScheduleForDay(client.NotionClient, time.Now())
			if err != nil {
				return tgbotapi.MessageConfig{}, err
			}
			msg = tgbotapi.NewMessage(chatId, answer)
			// case "Уведомлять о предстоящих занятиях":
			// 	if err := notion.SetScheduleNotifications(client.NotionClient, scheduler, schedulerTask); err != nil {
			// 		return tgbotapi.MessageConfig{}, err
			// 	}
			// 	msg = tgbotapi.NewMessage(chatId, "Уведомления активны")
		default:
			msg = tgbotapi.NewMessage(chatId, "Выберите из 3 вариантов!")
		}
	}
	return msg, nil

}

func SetStep(chatId int64, client NotionClientWithSteps, step Step) {
	client.CurrentStep = step
	clients[chatId] = client
}
