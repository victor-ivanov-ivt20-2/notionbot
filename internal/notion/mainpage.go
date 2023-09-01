package notion

import (
	"context"
	"errors"

	"github.com/jomei/notionapi"
)

func GetSchedule(client NotionClient) (*notionapi.TableBlock, error) {
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
	schedule, err := GetSchedule(client)

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
			SubjectRow, err := GetSubjectsToTableRow(index, client)

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

func GetSubjectsToTableRow(index int, client NotionClient) (notionapi.TableRow, error) {

	items, err := GetScheduleItems(client.Notion, client.ScheduleId, client.UserId, index)

	if err != nil {
		return notionapi.TableRow{}, err
	}

	scheduleRowCells := [][]notionapi.RichText{
		{
			{
				Text: &notionapi.Text{
					Content: "",
				},
			},
		},
		{
			{
				Text: &notionapi.Text{
					Content: "",
				},
			},
		},
		{
			{
				Text: &notionapi.Text{
					Content: "",
				},
			},
		},
		{
			{
				Text: &notionapi.Text{
					Content: "",
				},
			},
		},
		{
			{
				Text: &notionapi.Text{
					Content: "",
				},
			},
		},
		{
			{
				Text: &notionapi.Text{
					Content: "",
				},
			},
		},
	}

	WeekDays := GetWeekDays()
	EvenOdd := GetEvenOdd()
	LessonType := GetLessonType()

	for _, v := range items.Results {
		var title string
		var room string
		lessonType := -1
		evenodd := -1
		weekDay := -1
		for _, f := range v.Properties {
			if t, ok := f.(*notionapi.TitleProperty); ok {
				title = t.Title[0].Text.Content
			}
			if q, ok := f.(*notionapi.SelectProperty); ok {
				d, okd := WeekDays[q.Select.Name]
				e, oke := EvenOdd[q.Select.Name]
				l, okl := LessonType[q.Select.Name]
				if okd {
					weekDay = d
				} else if oke {
					evenodd = e
				} else if okl {
					lessonType = l
				}
			}
			if q, ok := f.(*notionapi.RichTextProperty); ok {
				room = q.RichText[0].Text.Content
			}
		}

		if len(title) != 0 && weekDay != -1 && evenodd != -1 && lessonType != -1 && len(room) != 0 {
			row := scheduleRowCells[weekDay][0]
			var eo = ""
			var lt = ""
			if evenodd == 1 {
				eo = "*"
			} else if evenodd == 2 {
				eo = "**"
			}
			if lessonType == 0 {
				lt = "лек"
			} else if lessonType == 1 {
				lt = "лаб"
			} else if lessonType == 2 {
				lt = "пр"
			}

			row.Text.Content = row.Text.Content + title + eo + " " + room + "[" + lt + "] "
			scheduleRowCells[weekDay][0] = row
		}
	}

	scheduleRow := notionapi.TableRow{
		Cells: scheduleRowCells,
	}

	return scheduleRow, nil
}
