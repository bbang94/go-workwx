package workwx_test

import (
	"github.com/bbang94/go-workwx"
	"net/http"
)

const (
	corpID     = "your_corpid"
	corpSecret = "your_corpsecret"
)

func ExampleWorkwx() {
	agentID := int64(1234567)

	client := workwx.New(corpID)

	// there're advanced options
	_ = workwx.New(
		corpID,
		workwx.WithQYAPIHost("http://localhost:8888"),
		workwx.WithHTTPClient(&http.Client{}),
	)

	// work with individual apps
	app := client.WithApp(corpSecret, agentID)
	app.SpawnAccessTokenRefresher()

	// see other examples for more details
}

func ExampleWorkwxApp_SendTextMessage() {
	agentID := int64(1234567)

	client := workwx.New(corpID)

	app := client.WithApp(corpSecret, agentID)
	// preferably do this at app initialization
	app.SpawnAccessTokenRefresher()

	// send to user(s)
	to1 := workwx.Recipient{
		UserIDs: []string{"testuser"},
	}
	_ = app.SendTextMessage(&to1, "send to user(s)", false)

	// "safe" message
	to2 := workwx.Recipient{
		UserIDs: []string{"testuser"},
	}
	_ = app.SendTextMessage(&to2, "safe message", true)

	// send to party(parties)
	to3 := workwx.Recipient{
		PartyIDs: []string{"testdept"},
	}
	_ = app.SendTextMessage(&to3, "send to party(parties)", false)

	// send to tag(s)
	to4 := workwx.Recipient{
		TagIDs: []string{"testtag"},
	}
	_ = app.SendTextMessage(&to4, "send to tag(s)", false)

	// send to chatid
	to5 := workwx.Recipient{
		ChatID: "testchat",
	}
	_ = app.SendTextMessage(&to5, "send to chatid", false)
}

func ExampleWorkwxApp_ApplyOAEvent() {
	agentID := int64(1234567)

	client := workwx.New(corpID)

	app := client.WithApp(corpSecret, agentID)
	app.SpawnAccessTokenRefresher()

	appInfo := workwx.OAApplyEvent{
		CreatorUserID:       "your_userid",
		TemplateID:          "your_templateid",
		UseTemplateApprover: 1,
		ApplyData: workwx.OAContents{
			Contents: []workwx.OAContent{
				{
					Control: workwx.OAControlText,
					ID:      "Text-1608628829793",
					Value: workwx.OAContentValue{
						Text: "文本",
					},
				},
				{
					Control: workwx.OAControlTextarea,
					ID:      "Textarea-1608628832640",
					Value: workwx.OAContentValue{
						Text: "多行文本\n可换行",
					},
				},
				{
					Control: workwx.OAControlNumber,
					ID:      "Number-1608632495498",
					Value: workwx.OAContentValue{
						Number: "123.45",
					},
				},
				{
					Control: workwx.OAControlMoney,
					ID:      "Money-1608632497034",
					Value: workwx.OAContentValue{
						Money: "678.90",
					},
				},
				{
					Control: workwx.OAControlContact,
					ID:      "Contact-1608632505579",
					Value: workwx.OAContentValue{
						Members: []workwx.OAContentMember{{
							UserID: "your_userid",
							Name:   "your_name",
						}},
					},
				},
				{
					Control: workwx.OAControlContact,
					ID:      "Contact-1608632506635",
					Value: workwx.OAContentValue{
						Departments: []workwx.OAContentDepartment{{
							OpenAPIID: "39",
							Name:      "xx部门1",
						}, {
							OpenAPIID: "40",
							Name:      "xx部门2",
						}},
					},
				},

				{
					Control: workwx.OAControlRelatedApproval,
					ID:      "RelatedApproval-1608632509930",
					Value: workwx.OAContentValue{
						RelatedApproval: []workwx.OAContentRelatedApproval{
							{SpNo: "202012220021"},
						},
					},
				},
				{
					Control: workwx.OAControlTable,
					ID:      "Table-1608632511066",
					Value: workwx.OAContentValue{
						Table: []workwx.OAContentTableList{
							{
								List: []workwx.OAContent{
									{
										Control: workwx.OAControlText,
										ID:      "Text-1608632519610",
										Value: workwx.OAContentValue{
											Text: "第一行第一列",
										},
									}, {
										Control: workwx.OAControlText,
										ID:      "Text-1608632521106",
										Value: workwx.OAContentValue{
											Text: "第一行第二列",
										},
									},
								},
							},
							{
								List: []workwx.OAContent{
									{
										Control: workwx.OAControlText,
										ID:      "Text-1608632519610",
										Value: workwx.OAContentValue{
											Text: "第二行第一列",
										},
									}, {
										Control: workwx.OAControlText,
										ID:      "Text-1608632521106",
										Value: workwx.OAContentValue{
											Text: "第二行第二列",
										},
									},
								},
							},
						},
					},
				},
			},
		},
		SummaryList: []workwx.OASummaryList{{SummaryInfo: []workwx.OAText{{
			Text: "摘要第1行",
		}}}, {SummaryInfo: []workwx.OAText{{
			Text: "摘要第2行",
		}}}},
	}
	_, _ = app.ApplyOAEvent(appInfo)
}
