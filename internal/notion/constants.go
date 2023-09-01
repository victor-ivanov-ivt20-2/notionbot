package notion

import "github.com/jomei/notionapi"

var HeaderTableRow = &notionapi.TableRow{
	Cells: [][]notionapi.RichText{
		{
			{
				Text: &notionapi.Text{
					Content: "7 семестр",
				},
			},
		},
		{
			{
				Text: &notionapi.Text{
					Content: "Понедельник",
				},
			},
		},
		{
			{
				Text: &notionapi.Text{
					Content: "Вторник",
				},
			},
		},
		{
			{
				Text: &notionapi.Text{
					Content: "Среда",
				},
			},
		},
		{
			{
				Text: &notionapi.Text{
					Content: "Четверг",
				},
			},
		},
		{
			{
				Text: &notionapi.Text{
					Content: "Пятница",
				},
			},
		},
		{
			{
				Text: &notionapi.Text{
					Content: "Суббота",
				},
			},
		},
	},
}

var TimeTableRow = [][]notionapi.RichText{
	{
		{
			Text: &notionapi.Text{
				Content: "08:00 - 09:35",
			},
		},
	},
	{
		{
			Text: &notionapi.Text{
				Content: "09:50 - 11:25",
			},
		},
	},
	{
		{
			Text: &notionapi.Text{
				Content: "11:40 - 13:15",
			},
		},
	},
	{
		{
			Text: &notionapi.Text{
				Content: "14:00 - 15:35",
			},
		},
	},
	{
		{
			Text: &notionapi.Text{
				Content: "15:50 - 17:25",
			},
		},
	},
	{
		{
			Text: &notionapi.Text{
				Content: "17:40 - 19:15",
			},
		},
	},
}

func GetWeekDays() map[string]int {
	var WeekDays = make(map[string]int)

	WeekDays["Понедельник"] = 0
	WeekDays["Вторник"] = 1
	WeekDays["Среда"] = 2
	WeekDays["Четверг"] = 3
	WeekDays["Пятница"] = 4
	WeekDays["Суббота"] = 5
	return WeekDays
}
func GetEvenOdd() map[string]int {
	var EvenOdd = make(map[string]int)

	EvenOdd["чётное/нечётное"] = 0
	EvenOdd["нечётное"] = 1
	EvenOdd["чётное"] = 2
	return EvenOdd
}
func GetLessonType() map[string]int {
	var LessonType = make(map[string]int)

	LessonType["лекция"] = 0
	LessonType["лабораторная"] = 1
	LessonType["практика"] = 2
	return LessonType
}
