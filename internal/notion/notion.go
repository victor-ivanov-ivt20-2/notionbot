package notion

import (
	"unicode"

	"github.com/jomei/notionapi"
)

type NotionClient struct {
	Notion     *notionapi.Client
	ScheduleId string
	TasksId    string
	UserId     string
	PageId     string
	Email      string
}

var WeekDays = GetWeekDays()
var EvenOdd = GetEvenOdd()
var LessonType = GetLessonType()
var LessonTime = GetLessonTime()

func SetClient(token notionapi.Token) *notionapi.Client {
	client := notionapi.NewClient(token)
	return client
}

func RefreshSchedule() {

}

func checkRoom(str string) string {
	for _, ch := range str {
		if unicode.IsDigit(ch) {
			return " в " + str + " кабинете"
		}
	}
	return str
}

func CheckAllProperties(withoutValue bool, args ...interface{}) bool {

	for _, v := range args {
		switch t := v.(type) {
		case string:
			if withoutValue {
				if len(t) == 0 {
					return false
				}
			} else {
				if t == "-1" {
					return false
				}
			}

		case int:
			if t == -1 {
				return false
			}
		}
	}
	return true

}

// func CreateBlock(notion *notionapi.Client, pageId string, text string) (*notionapi.AppendBlockChildrenResponse, error) {
// 	createdBlock, err := notion.Block.AppendChildren(context.Background(), notionapi.BlockID(pageId), &notionapi.AppendBlockChildrenRequest{
// 		Children: []notionapi.Block{
// 			&notionapi.ParagraphBlock{
// 				BasicBlock: notionapi.BasicBlock{
// 					Object: "block",
// 					Type:   "paragraph",
// 				},
// 				Paragraph: notionapi.Paragraph{
// 					RichText: []notionapi.RichText{
// 						{
// 							Type: "",
// 							Text: &notionapi.Text{
// 								Content: text,
// 							},
// 							PlainText: text,
// 						},
// 					},
// 					Color: "gray",
// 				},
// 			},
// 		},
// 	})
// 	if err != nil {
// 		return nil, errors.New(err.Error())
// 	}
// 	return createdBlock, nil
// }

// func extractPlainTextTitle(title []notionapi.RichText) string {
// 	var plainTextTitle string
// 	for _, richText := range title {
// 		plainTextTitle += richText.PlainText
// 	}
// 	return plainTextTitle
// }
