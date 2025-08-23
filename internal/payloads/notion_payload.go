package notion_payload

import "time"

type NotifyNHKTaskPayload struct {
	Event Event `json:"event"`
}

type Event struct {
	Type string `json:"type"`
}

type Page struct {
	Id             string    `json:"id"`
	LastEditedTime time.Time `json:"last_edited_time"`
}

type PageProperties struct {
	Name   string `json:"Name"`
	Status string `json:"status"`
}
