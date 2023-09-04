package bot

import (
	"errors"

	"github.com/go-co-op/gocron"
	"github.com/victor-ivanov-ivt20-2/ourdiary/internal/notion"
)

type NotionClientWithSteps struct {
	CurrentStep  Step
	Scheduler    *gocron.Scheduler
	NotionClient notion.NotionClient
}

func GetClient(chatId int64) (NotionClientWithSteps, error) {
	client, ok := clients[chatId]
	// This condition initializes the client
	if !ok {
		clients[chatId] = NotionClientWithSteps{CurrentStep: WELCOME}
		client, ok = clients[chatId]
		// This condition checks for an error
		if !ok {
			return NotionClientWithSteps{}, errors.New("failed to init client")
		}
	}

	return client, nil
}
