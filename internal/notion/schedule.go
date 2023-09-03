package notion

import (
	"time"

	"github.com/go-co-op/gocron"
	"github.com/jomei/notionapi"
	"github.com/victor-ivanov-ivt20-2/ourdiary/internal/lib/scheduler"
)

type SchedulerTask func(string, string, string) error

func GetScheduleForDay(client NotionClient, t time.Time) (string, error) {
	tomorrow := scheduler.GetNextWeekDay(t)
	if tomorrow == "Воскресенье" {
		return "Выходной день", nil
	}
	var evenodd string
	if scheduler.GetEvenOddWeek(t) {
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

func SetScheduleNotifications(client NotionClient, scheduler *gocron.Scheduler, schedulerTask SchedulerTask) error {

	// items, err := GetScheduleItems(client.Notion, client.ScheduleId, client.UserId,
	// 	notionapi.PropertyFilter{
	// 		Property: "Person",
	// 		People: &notionapi.PeopleFilterCondition{
	// 			Contains: client.UserId,
	// 		},
	// 	},
	// 	[]notionapi.SortObject{})

	// if err != nil {
	// 	return err
	// }

	// for _, r := range items.Results {
	// 	for _, p := range r.Properties {
	// 		var lessonTimeStart string
	// 		var lessonTimeEnd string
	// 		var weekDay = -1
	// 		var title string
	// 		var room string
	// 		if t, ok := p.(*notionapi.TitleProperty); ok {
	// 			title = t.Title[0].Text.Content
	// 		}
	// 		if n, ok := p.(*notionapi.NumberProperty); ok {
	// 			lessonTimeStart = LessonTime[int(n.Number)][0:4]
	// 			lessonTimeEnd = LessonTime[int(n.Number)][8:12]
	// 		}
	// 		if w, ok := p.(*notionapi.SelectProperty); ok {
	// 			weekDay = WeekDays[w.Select.Name]
	// 		}
	// 		if r, ok := p.(*notionapi.RichTextProperty); ok {
	// 			room = r.RichText[0].Text.Content
	// 		}

	// 		switch weekDay {
	// 		case 0:
	// _, errStart := scheduler.Every(1).Monday().At(lessonTimeStart).Do(schedulerTask, title, lessonTimeStart, room)
	// _, errStart := scheduler.Every(15).Seconds().Do(schedulerTask, title, lessonTimeStart, room)
	// if errStart != nil {
	// 	return errStart
	// }
	// _, errEnd := scheduler.Every(1).Monday().At(lessonTimeEnd).Do(schedulerTask, title, lessonTimeEnd, room)
	// _, errEnd := scheduler.Every(2).Seconds().Do(schedulerTask, title, lessonTimeStart, room)
	// fmt.Print(lessonTimeEnd)
	// if errEnd != nil {
	// 	return errEnd
	// }
	// 	}
	// }
	// }

	return nil
}
