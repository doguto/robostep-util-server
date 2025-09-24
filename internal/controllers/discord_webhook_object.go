package notion_controller

type DiscordWebhookObject struct {
	Embeds []Embed `json:"embeds"`
}

type Embed struct {
	Title       string  `json:"title"`
	URL         string  `json:"url"`
	Description string  `json:"description,omitempty"`
	Fields      []Field `json:"fields"`
	Color       int     `json:"color"`
}

type Field struct {
	Name   string `json:"name"`
	Value  string `json:"value"`
	Inline bool   `json:"inline"`
}
