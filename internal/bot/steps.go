package bot

import (
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
		tgbotapi.NewKeyboardButton("Расписание на завтра"),
	),
)

func Steps(chatId int64, message string, env config.OurDiary) (tgbotapi.MessageConfig, error) {
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

		switch message {
		case "Обновить расписание":
			if err := notion.UpdateSchedule(client.NotionClient); err != nil {
				return tgbotapi.MessageConfig{}, err
			}
			msg = tgbotapi.NewMessage(chatId, "Расписание на вашей странице было обновлено!")
		case "Расписание на завтра":
			answer, err := notion.GetScheduleTommorow(client.NotionClient)
			if err != nil {
				return tgbotapi.MessageConfig{}, err
			}
			msg = tgbotapi.NewMessage(chatId, answer)
		}
	}

	return msg, nil
}

func SetStep(chatId int64, client NotionClientWithSteps, step Step) {
	client.CurrentStep = step
	clients[chatId] = client
}
