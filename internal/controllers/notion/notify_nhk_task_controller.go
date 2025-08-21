package notion_controller

import (
	"github.com/gin-gonic/gin"
)

type NotionController struct {
}

func NewNotionController() *NotionController {
	return &NotionController{}
}

func (c *NotionController) NotifyTaskToDiscord(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "Task notification sent to Discord",
	})
}
