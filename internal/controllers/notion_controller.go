package notion_controller

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"os"

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

	// Discordへ通知
	payload := map[string]string{
		"content": "Test",
	}
	jsonPayload, _ := json.Marshal(payload)

	response, error := http.Post(
		webhookUrl,
		"application/json",
		bytes.NewBuffer(jsonPayload),
	)
	if error != nil {
		ctx.JSON(500, gin.H{
			"error": "Failed to send task notification to Discord",
		})
		return
	}
	defer response.Body.Close()

	body, _ := io.ReadAll(response.Body)

	ctx.JSON(200, gin.H{
		"message": "Task notification sent to Discord",
		"payload": body,
	})
}
