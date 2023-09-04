package notion

import (
	"time"

	"github.com/jomei/notionapi"
	"github.com/victor-ivanov-ivt20-2/ourdiary/internal/lib/scheduler"
)

type SchedulerTask func(string, string, string) error

func GetAllSchedule(client NotionClient) (string, error) {

	items, err := GetScheduleItems(client.Notion, client.ScheduleId, client.UserId, notionapi.PropertyFilter{
		Property: "Person",
		People: &notionapi.PeopleFilterCondition{
			Contains: client.UserId,
		},
	}, []notionapi.SortObject{
		{
			Property:  "StartDay",
			Direction: "ascending",
			Timestamp: "1",
		},
		{
			Property:  "StartTime",
			Direction: "ascending",
			Timestamp: "1",
		},
	})

	if err != nil {
		return "", nil
	}

	var answer string

	if scheduler.GetEvenOddWeek(time.Now()) {
		answer = "чётная неделя\n"
	} else {
		answer = "нечётная неделя\n"
	}

	var weekday = -1
	for _, v := range items.Results {
		var title = "-1"
		var evenodd = "-1"
		var lessonTime = "-1"
		var room = "-1"
		var lessonType = "-1"
		var weekDay = -1
		var weekDayString string
		for _, p := range v.Properties {
			if t, ok := p.(*notionapi.TitleProperty); ok {
				title = t.Title[0].Text.Content
			}
			if n, ok := p.(*notionapi.NumberProperty); ok {
				lessonTime = LessonTime[int(n.Number)]
			}
			if q, ok := p.(*notionapi.SelectProperty); ok {
				if d, okd := WeekDays[q.Select.Name]; okd {
					weekDay = d
					weekDayString = q.Select.Name
				}
				if e, okd := EvenOdd[q.Select.Name]; okd {
					evenodd = e
				}
				if l, okd := LessonType[q.Select.Name]; okd {
					lessonType = l
				}
			}
			if q, ok := p.(*notionapi.RichTextProperty); ok {
				room = q.RichText[0].Text.Content
			}
		}

		if weekDay != weekday {
			answer = answer + "\n" + weekDayString + "\n"
			weekday = weekDay
		}
		if CheckAllProperties(false, title, evenodd, lessonTime, room, lessonType, weekDay) {
			total := lessonTime + " : " + title + " " + evenodd + " " + room + " [" + lessonType + "] " + "\n"
			answer += total
		}
	}
	return answer, nil
}

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

func SetScheduleNotifications(client NotionClient) ([][]string, error) {

	items, err := GetScheduleItems(client.Notion, client.ScheduleId, client.UserId,
		notionapi.PropertyFilter{
			Property: "Person",
			People: &notionapi.PeopleFilterCondition{
				Contains: client.UserId,
			},
		},
		[]notionapi.SortObject{
			{
				Property:  "StartDay",
				Direction: "ascending",
				Timestamp: "1",
			},
			{
				Property:  "StartTime",
				Direction: "ascending",
				Timestamp: "1",
			},
		})

	if err != nil {
		return nil, err
	}
	var response [][]string

	for j := 0; j < 6; j++ {
		response = append(response, []string{})
	}

	for _, r := range items.Results {
		var lessonTimeStart string
		var lessonTimeEnd string
		var weekDay = -1
		var title string
		var evenodd string
		var room string
		for _, p := range r.Properties {
			if t, ok := p.(*notionapi.TitleProperty); ok {
				title = t.Title[0].Text.Content
			}
			if n, ok := p.(*notionapi.NumberProperty); ok {
				lessonTimeStart = LessonTime[int(n.Number)][0:5]
				lessonTimeEnd = LessonTime[int(n.Number)][8:13]
			}
			if w, ok := p.(*notionapi.SelectProperty); ok {
				if d, okd := WeekDays[w.Select.Name]; okd {
					weekDay = d
				}
				if e, okd := EvenOdd[w.Select.Name]; okd {
					evenodd = e
				}
			}
			if r, ok := p.(*notionapi.RichTextProperty); ok {
				room = r.RichText[0].Text.Content
			}
		}

		if !CheckAllProperties(true, lessonTimeStart, lessonTimeEnd, weekDay, title, room) {
			continue
		}

		if scheduler.GetEvenOddWeek(time.Now()) {
			if evenodd == "*" {
				continue
			}
		} else {
			if evenodd == "**" {
				continue
			}
		}

		if len(response[weekDay]) == 0 {
			response[weekDay] = append(response[weekDay], lessonTimeStart)
			response[weekDay] = append(response[weekDay], title+" начнётся в "+lessonTimeStart+checkRoom(room))
			response[weekDay] = append(response[weekDay], lessonTimeEnd)
		} else {
			response[weekDay] = append(response[weekDay], "Следующая пара "+title+" начнётся в "+lessonTimeStart+checkRoom(room))
			response[weekDay] = append(response[weekDay], lessonTimeEnd)
		}
	}
	// switch weekDay {
	// case 0:
	// _, errStart := scheduler.Every(1).Monday().At(lessonTimeStart).Do(schedulerTask, title, lessonTimeStart, room)
	// _, errEnd := scheduler.Every(1).Monday().At(lessonTimeEnd).Do(schedulerTask, title, lessonTimeEnd, room)
	// }

	return response, nil
}
