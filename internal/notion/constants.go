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
func GetEvenOdd() map[string]string {
	var EvenOdd = make(map[string]string)

	EvenOdd["чётное/нечётное"] = ""
	EvenOdd["нечётное"] = "*"
	EvenOdd["чётное"] = "**"
	return EvenOdd
}
func GetLessonType() map[string]string {
	var LessonType = make(map[string]string)

	LessonType["лекция"] = "лек"
	LessonType["лабораторная"] = "лаб"
	LessonType["практика"] = "пр"
	return LessonType
}

func GetLessonTime() map[int]string {
	var LessonTime = make(map[int]string)

	LessonTime[1] = "08:00 - 09:35"
	LessonTime[2] = "09:50 - 11:25"
	LessonTime[3] = "11:40 - 13:15"
	LessonTime[4] = "14:00 - 15:35"
	LessonTime[5] = "15:50 - 17:25"
	LessonTime[6] = "17:40 - 19:15"

	return LessonTime
}

// func getKeyFromWeekDay(m map[string]int, value int) string {
// 	for k, v := range m {
// 		if v == value {
// 			return k
// 		}
// 	}
// 	return ""
// }
