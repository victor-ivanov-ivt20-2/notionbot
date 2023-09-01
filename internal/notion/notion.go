package notion

import (
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

func SetClient(token notionapi.Token) *notionapi.Client {
	client := notionapi.NewClient(token)
	return client
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
