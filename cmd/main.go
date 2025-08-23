package main

import (
	"github.com/gin-gonic/gin"

	notion_controller "robostep-util-server/internal/controllers"
)

func main() {
	engine := gin.Default()

	notionController := notion_controller.NewNotionController()

	// API routing
	apiGroup := engine.Group("/api")
	{
		notionGroup := apiGroup.Group("/notion")
		{
			notionGroup.GET("/notify_nhk_task", notionController.NotifyTaskToDiscord)
		}
	}

	engine.Run(":8080")
}
