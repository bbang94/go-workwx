/*
* Author:  dqx
* Date:    2024/3/27 14:23
* contact: 769989770@qq.com
* Description:
 */

package workwx

import (
	"encoding/json"
	"testing"
)

func TestOAText(t *testing.T) {
	oaContens := []OAContent{
		{
			Control: "Text",
			ID:      "Text-1629290443705",
			Value:   OAContentValue{Text: "2312"},
		},
	}
	applyInfo := OAContents{Contents: oaContens}
	event := OAApplyEvent{
		CreatorUserID:       "user.VxId",
		TemplateID:          "C4WqcQGkRmsVrdyS3grZFtNDZpDhCHiWEFQBxB9qu",
		UseTemplateApprover: 1,
		ApplyData:           applyInfo,
	}
	marshal, err := json.Marshal(event)
	if err != nil {
		return
	}
	t.Log(marshal)
}
