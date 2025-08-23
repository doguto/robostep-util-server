package notion_payload

// NotifyNhkTaskPayload は Notion Webhook から受け取るページ変更ペイロードの構造体
type NotifyNhkTaskPayload struct {
	Source struct {
		Type         string `json:"type"`
		AutomationId string `json:"automation_id"`
		ActionId     string `json:"action_id"`
		EventId      string `json:"event_id"`
		Attempt      int    `json:"attempt"`
	} `json:"source"`
	Data struct {
		Object         string `json:"object"`
		Id             string `json:"id"`
		CreatedTime    string `json:"created_time"`
		LastEditedTime string `json:"last_edited_time"`
		CreatedBy      struct {
			Object string `json:"object"`
			Id     string `json:"id"`
		} `json:"created_by"`
		LastEditedBy struct {
			Object string `json:"object"`
			Id     string `json:"id"`
		} `json:"last_edited_by"`
		Cover string `json:"cover"`
		Icon  struct {
			Type     string `json:"type"`
			External struct {
				Url string `json:"url"`
			} `json:"external"`
		} `json:"icon"`
		Parent struct {
			Type       string `json:"type"`
			DatabaseId string `json:"database_id"`
		} `json:"parent"`
		Archived bool `json:"archived"`
		InTrash  bool `json:"in_trash"`
	} `json:"data"`
	Properties DatabaseProperty `json:"properties"` // データベースプロパティをここに格納
	URL        string           `json:"url"`
	PublicUrl  string           `json:"public_url"`
	RequestId  string           `json:"request_id"`
}

// DatabaseProperty は Notion データベースの各プロパティを定義
type DatabaseProperty struct {
	Status struct {
		Id     string `json:"id"`
		Type   string `json:"type"`
		Status struct {
			Id    string `json:"id"`
			Name  string `json:"name"`
			Color string `json:"color"`
		} `json:"status"`
	} `json:"ステータス"`

	ParentTask struct {
		Id       string `json:"id"`
		Type     string `json:"type"`
		Relation []struct {
			Id string `json:"id"`
		} `json:"relation"`
		HasMore bool `json:"has_more"`
	} `json:"親タスク"`

	TaskKind struct {
		Id          string `json:"id"`
		Type        string `json:"type"`
		MultiSelect []struct {
			Id    string `json:"id"`
			Name  string `json:"name"`
			Color string `json:"color"`
		} `json:"multi_select"`
	} `json:"タスクの種類"`

	Cost struct {
		Id     string      `json:"id"`
		Type   string      `json:"type"`
		Select interface{} `json:"select"` // TODO: 確認次第型を定義
	} `json:"工数レベル"`

	Limit struct {
		Id   string      `json:"id"`
		Type string      `json:"type"`
		Date interface{} `json:"date"` // TODO: 確認次第型を定義
	} `json:"期日"`

	Assignees struct {
		Id     string `json:"id"`
		Type   string `json:"type"`
		People []struct {
			Object    string `json:"object"`
			Id        string `json:"id"`
			Name      string `json:"name"`
			AvatarURL string `json:"avatar_url"`
			Type      string `json:"type"`
			Person    struct {
				Email string `json:"email"`
			} `json:"person"`
		} `json:"people"`
	} `json:"担当者"`

	SubTasks struct {
		Id       string `json:"id"`
		Type     string `json:"type"`
		Relation []any  `json:"relation"`
		HasMore  bool   `json:"has_more"`
	} `json:"サブタスク"`

	TaskName struct {
		Id    string `json:"id"`
		Type  string `json:"type"`
		Title []struct {
			Type string `json:"type"`
			Text struct {
				Content string `json:"content"`
				Link    any    `json:"link"`
			} `json:"text"`
			Annotations struct {
				Bold          bool   `json:"bold"`
				Italic        bool   `json:"italic"`
				Strikethrough bool   `json:"strikethrough"`
				Underline     bool   `json:"underline"`
				Code          bool   `json:"code"`
				Color         string `json:"color"`
			} `json:"annotations"`
			PlainText string `json:"plain_text"`
			Href      any    `json:"href"`
		} `json:"title"`
	} `json:"タスク名"`
}
