package notion

import (
	"context"

	"github.com/jomei/notionapi"
)

func GetScheduleItems(notion *notionapi.Client, scheduleId string, userId string, index int) (*notionapi.DatabaseQueryResponse, error) {
	floatIndex := float64(index)
	databaseItems, err := notion.Database.Query(context.Background(), notionapi.DatabaseID(scheduleId), &notionapi.DatabaseQueryRequest{
		Filter: notionapi.AndCompoundFilter{
			notionapi.PropertyFilter{
				Property: "Person",
				People: &notionapi.PeopleFilterCondition{
					Contains: userId,
				},
			},
			notionapi.PropertyFilter{
				Property: "StartTime",
				Number: &notionapi.NumberFilterCondition{
					Equals: &floatIndex,
				},
			},
		},
	})

	if err != nil {
		return nil, err
	}

	return databaseItems, nil

}

// func GetDatabaseItems(notion *notionapi.Client, databaseId string) (*notionapi.DatabaseQueryResponse, error) {

// 	databaseItems, err := notion.Database.Query(context.Background(), notionapi.DatabaseID(databaseId), &notionapi.DatabaseQueryRequest{})
// 	if err != nil {
// 		return nil, err
// 	}

// 	for _, v := range databaseItems.Results {
// 		for _, f := range v.Properties {
// 			if q, ok := f.(*notionapi.NumberProperty); ok {
// 				fmt.Print(q.Number)
// 			}
// 			if q, ok := f.(*notionapi.SelectProperty); ok {
// 				fmt.Print(q.Select.Name)
// 			}
// 			fmt.Print(" ")
// 		}
// 		fmt.Println("")
// 	}

// 	return databaseItems, nil
// }

// func AddDatabaseItem(notion *notionapi.Client, env config.OurDiary) error {
// 	_, err := notion.Page.Create(context.Background(), &notionapi.PageCreateRequest{
// 		Parent: notionapi.Parent{
// 			DatabaseID: notionapi.DatabaseID(env.ScheduleId),
// 		},
// 		Properties: notionapi.Properties{
// 			"title": notionapi.TitleProperty{
// 				Title: []notionapi.RichText{
// 					{
// 						Text: &notionapi.Text{
// 							Content: "new Item",
// 						},
// 					},
// 				},
// 			},
// 			"Person": notionapi.PeopleProperty{

// 				People: []notionapi.User{
// 					{
// 						ID:   notionapi.UserID(env.Second.UserId),
// 						Type: "person",
// 						Person: &notionapi.Person{
// 							Email: env.Second.Email,
// 						},
// 						Bot: nil,
// 					},
// 				},
// 			},
// 		},
// 	})
// 	return err
// }
