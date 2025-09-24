package notion_controller

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	notion_payload "robostep-util-server/internal/payloads"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

type NotionController struct {
}

func NewNotionController() *NotionController {
	return &NotionController{}
}

func (c *NotionController) NotifyTaskToDiscord(ctx *gin.Context) {
	// 環境変数の設定
	godotenv.Load()
	webhookUrl := os.Getenv("DISCORD_WEBHOOK_URL")

	// リクエストボディの取得
	rawData, _ := ctx.GetRawData()
	var payload notion_payload.NotifyNhkTaskPayload
	err := json.Unmarshal(rawData, &payload)
	if err != nil {
		fmt.Printf("JSON Unmarshal Error: %v\n", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	var taskName string
	if len(payload.Data.Properties.TaskName.Title) > 0 {
		taskName = payload.Data.Properties.TaskName.Title[0].PlainText
	} else {
		taskName = "（タスク名なし）"
	}

	var assignees string
	if len(payload.Data.Properties.Assignees.People) > 0 {
		for i, a := range payload.Data.Properties.Assignees.People {
			if i > 0 {
				assignees += ", "
			}
			assignees += a.Name
		}
	} else {
		assignees = "（担当者なし）"
	}

	// Discordへ通知
	var limitDate string
	if payload.Data.Properties.Limit.Date.Start != "" {
		limitDate = payload.Data.Properties.Limit.Date.Start
	} else {
		limitDate = "（期日なし）"
	}

	var status = payload.Data.Properties.Status.Status.Name

	var taskKinds = payload.Data.Properties.TaskKind.MultiSelect
	var taskKind = ""
	for index, multiSelect := range taskKinds {
		var taskKindName = multiSelect.Name
		if index != 0 {
			taskKind += ", "
		}
		taskKind += taskKindName
	}

	var discordWebhookObject DiscordWebhookObject
	discordWebhookObject.Embeds = []Embed{
		{
			Title: "タスクリストが更新されました！",
			URL:   payload.Data.URL,
			Fields: []Field{
				{
					Name:   "タスク名",
					Value:  taskName,
					Inline: false,
				},
				{
					Name:   "期日",
					Value:  limitDate,
					Inline: false,
				},
				{
					Name:   "担当者",
					Value:  assignees,
					Inline: false,
				},
				{
					Name:   "ステータス",
					Value:  status,
					Inline: true,
				},
				{
					Name:   "タスクの種類",
					Value:  taskKind,
					Inline: true,
				},
			},
			Color: 5763719, // Green
		},
	}

	jsonBody, _ := json.Marshal(discordWebhookObject)

	response, error := http.Post(
		webhookUrl,
		"application/json",
		bytes.NewBuffer(jsonBody),
	)
	if error != nil {
		ctx.JSON(500, gin.H{
			"error": "Failed to send task notification to Discord",
		})
		return
	}
	defer response.Body.Close()

	ctx.JSON(200, gin.H{
		"message": "Task notification sent to Discord",
		"payload": payload,
	})
}
