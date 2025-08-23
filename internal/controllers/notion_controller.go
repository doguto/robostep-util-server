package notion_controller

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
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

	var payload notion_payload.NotifyNhkTaskPayload
	err := ctx.ShouldBindJSON(&payload)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	body, _ := io.ReadAll(ctx.Request.Body)

	// Discordへ通知
	noticeBody := map[string]string{
		"content": fmt.Sprintf(`=== === ===
		### タスクリストが更新されました！
		　タスク名：**%s**
		　　担当者：**%s**
		ステータス：**%s**
		　　　期日：**%s**
		`,
			payload.Properties.TaskName.Title[0].PlainText,
			payload.Properties.Assignees.People[0].Name,
			payload.Properties.Status.Status.Name,
			payload.Properties.Limit.Date),
	}
	jsonBody, _ := json.Marshal(noticeBody)

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
		"payload": body,
	})
}
