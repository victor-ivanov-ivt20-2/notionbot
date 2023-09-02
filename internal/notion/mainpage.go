package notion

import (
	"context"
	"errors"

	"github.com/jomei/notionapi"
)

func GetScheduleTable(client NotionClient) (*notionapi.TableBlock, error) {
	schedule, err := client.Notion.Block.GetChildren(context.Background(), notionapi.BlockID(client.PageId), &notionapi.Pagination{StartCursor: notionapi.Cursor(""), PageSize: 10})

	if err != nil {
		return nil, err
	}

	for _, v := range schedule.Results {
		if q, ok := v.(*notionapi.TableBlock); ok {
			return q, nil
		}
	}

	return nil, errors.New("table not found")
}

func UpdateSchedule(client NotionClient) error {
	schedule, err := GetScheduleTable(client)

	if err != nil {
		return err
	}

	scheduleRows, err := client.Notion.Block.GetChildren(context.Background(), schedule.ID, &notionapi.Pagination{
		PageSize: 100,
	})

	if err != nil {
		return err
	}

	for index, c := range scheduleRows.Results {
		if q, ok := c.(*notionapi.TableRowBlock); ok {
			SubjectRow, err := SetSubjectsToTableRow(index, client)

			if err != nil {
				return err
			}

			if _, err := client.Notion.Block.Update(context.Background(), q.ID, &notionapi.BlockUpdateRequest{
				TableRow: GetUpdatedTableRow(index, SubjectRow),
			}); err != nil {
				return err
			}
		}
	}

	return nil
}

func GetUpdatedTableRow(index int, subjects notionapi.TableRow) *notionapi.TableRow {
	if index == 0 {
		return HeaderTableRow
	} else {
		var tableRow = &notionapi.TableRow{}
		tableRow.Cells = append(tableRow.Cells, TimeTableRow[index-1])
		tableRow.Cells = append(tableRow.Cells, subjects.Cells...)
		return tableRow
	}
}

func SetSubjectsToTableRow(index int, client NotionClient) (notionapi.TableRow, error) {
	floatIndex := float64(index)
	items, err := GetScheduleItems(client.Notion, client.ScheduleId, client.UserId, notionapi.AndCompoundFilter{
		notionapi.PropertyFilter{
			Property: "Person",
			People: &notionapi.PeopleFilterCondition{
				Contains: client.UserId,
			},
		},
		notionapi.PropertyFilter{
			Property: "StartTime",
			Number: &notionapi.NumberFilterCondition{
				Equals: &floatIndex,
			},
		},
	}, []notionapi.SortObject{})

	if err != nil {
		return notionapi.TableRow{}, err
	}

	var scheduleRowCells [][]notionapi.RichText
	for i := 0; i < 6; i++ {
		var item = []notionapi.RichText{}
		scheduleRowCells = append(scheduleRowCells, item)
	}

	for _, v := range items.Results {
		var title = "-1"
		var room = "-1"
		var lessonType = "-1"
		var evenodd = "-1"
		var url = v.URL
		weekDay := -1

		for _, f := range v.Properties {
			if t, ok := f.(*notionapi.TitleProperty); ok {
				title = t.Title[0].Text.Content
			}
			if q, ok := f.(*notionapi.SelectProperty); ok {
				if d, okd := WeekDays[q.Select.Name]; okd {
					weekDay = d
				}
				if e, okd := EvenOdd[q.Select.Name]; okd {
					evenodd = e
				}
				if l, okd := LessonType[q.Select.Name]; okd {
					lessonType = l
				}
			}
			if q, ok := f.(*notionapi.RichTextProperty); ok {
				room = q.RichText[0].Text.Content
			}
		}

		if weekDay != -1 && title != "-1" && evenodd != "-1" && lessonType != "-1" && room != "-1" {
			titleText := notionapi.RichText{
				Text: &notionapi.Text{
					Content: title,
					// TODO: Создать базу данных настоящих пар, куда эти ссылки и должны будут вести.
					// Вероятнее всего брать ссылку будут из properties
					Link: &notionapi.Link{
						Url: url,
					},
				},
				PlainText: title,
				Href:      url,
			}
			evenoddText := notionapi.RichText{
				Text: &notionapi.Text{
					Content: evenodd + " ",
				},
				PlainText: evenodd + " ",
			}
			roomText := notionapi.RichText{
				Text: &notionapi.Text{
					Content: room + " ",
				},
				PlainText: room + " ",
			}
			lessonTypeText := notionapi.RichText{
				Text: &notionapi.Text{
					Content: "[" + lessonType + "]\n",
				},
				PlainText: "[" + lessonType + "]",
			}
			var color notionapi.Color
			switch lessonType {
			case "лаб":
				color = notionapi.ColorBlueBackground
			case "лек":
				color = notionapi.ColorGreenBackground
			case "пр":
				color = notionapi.ColorRedBackground
			}
			lessonTypeText.Annotations = &notionapi.Annotations{
				Color: color,
			}
			scheduleRowCells[weekDay] = append(scheduleRowCells[weekDay], titleText)
			scheduleRowCells[weekDay] = append(scheduleRowCells[weekDay], evenoddText)
			scheduleRowCells[weekDay] = append(scheduleRowCells[weekDay], roomText)
			scheduleRowCells[weekDay] = append(scheduleRowCells[weekDay], lessonTypeText)
		}
	}

	scheduleRow := notionapi.TableRow{
		Cells: scheduleRowCells,
	}

	return scheduleRow, nil
}
