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

func Steps(chatId int64, message string, env config.OurDiary) (tgbotapi.MessageConfig, error) {
	var msg tgbotapi.MessageConfig
	client, err := GetClient(chatId)

	if err != nil {
		return tgbotapi.MessageConfig{}, err
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
		if message == env.First.Email {
			msg = tgbotapi.NewMessage(chatId, "Поздравляю!")
			client.NotionClient.TasksId = env.TasksId
			client.NotionClient.ScheduleId = env.ScheduleId
			client.NotionClient.Email = env.First.Email
			client.NotionClient.PageId = env.First.PageId
			client.NotionClient.UserId = env.First.UserId
			SetStep(chatId, client, WORKING)
		} else if message == env.Second.Email {
			msg = tgbotapi.NewMessage(chatId, "Поздравляю!")
			client.NotionClient.TasksId = env.TasksId
			client.NotionClient.ScheduleId = env.ScheduleId
			client.NotionClient.Email = env.Second.Email
			client.NotionClient.PageId = env.Second.PageId
			client.NotionClient.UserId = env.Second.UserId
			SetStep(chatId, client, WORKING)
		} else {
			msg = tgbotapi.NewMessage(chatId, "Повторяй")
		}
	case WORKING:
		msg = tgbotapi.NewMessage(chatId, "Ты теперь работаешь...")
		if err := notion.UpdateSchedule(client.NotionClient); err != nil {
			return tgbotapi.MessageConfig{}, err
		}
	}

	return msg, nil
}

func SetStep(chatId int64, client NotionClientWithSteps, step Step) {
	client.CurrentStep = step
	clients[chatId] = client
}
