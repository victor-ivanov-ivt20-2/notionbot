package notion

import (
	"time"

	"github.com/jomei/notionapi"
	"github.com/victor-ivanov-ivt20-2/ourdiary/internal/lib/gocron"
)

func GetScheduleTomorrow(client NotionClient) (string, error) {

	tomorrow := gocron.GetNextWeekDay()
	var evenodd string
	if gocron.GetEvenOddWeek(time.Now().AddDate(0, 0, 1)) {
		evenodd = "нечётное"
	} else {
		evenodd = "чётное"
	}
	items, err := GetScheduleItems(client.Notion, client.ScheduleId, client.UserId, notionapi.AndCompoundFilter{
		notionapi.PropertyFilter{
			Property: "StartDay",
			Select: &notionapi.SelectFilterCondition{
				Equals: tomorrow,
			},
		},
		notionapi.PropertyFilter{
			Property: "Status",
			Select: &notionapi.SelectFilterCondition{
				DoesNotEqual: evenodd,
			},
		},
		notionapi.PropertyFilter{
			Property: "Person",
			People: &notionapi.PeopleFilterCondition{
				Contains: client.UserId,
			},
		},
	}, []notionapi.SortObject{
		{
			Property:  "StartTime",
			Direction: "ascending",
			Timestamp: "1",
		},
	})

	if err != nil {
		return "", err
	}

	var answer string
	for _, v := range items.Results {
		var total string
		var title string
		var room string
		var number int
		for _, p := range v.Properties {
			if t, ok := p.(*notionapi.TitleProperty); ok {
				title = t.Title[0].Text.Content
			}
			if n, ok := p.(*notionapi.NumberProperty); ok {
				number = int(n.Number)
			}
			if q, ok := p.(*notionapi.RichTextProperty); ok {
				room = q.RichText[0].Text.Content
			}
		}
		total = LessonTime[number] + " : " + title + " " + room + "\n"
		answer += total
	}
	return answer, nil
}
