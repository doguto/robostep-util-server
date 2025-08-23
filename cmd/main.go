package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	notion_controller "robostep-util-server/internal/controllers"
)

func main() {
	engine := gin.Default()

	notionController := notion_controller.NewNotionController()

	// API routing
	apiGroup := engine.Group("/api")
	{
		apiGroup.GET("/", func(ctx *gin.Context) {
			ctx.String(http.StatusOK, "ok")
		})

		notionGroup := apiGroup.Group("/notion")
		{
			notionGroup.POST("/notify_nhk_task", notionController.NotifyTaskToDiscord)
		}
	}

	engine.Run(":8080")
}
