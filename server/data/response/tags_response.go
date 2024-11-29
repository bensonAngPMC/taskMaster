package response

type TagsResponse struct {
	ID              uint             `json:"id"`
	Name            string          `json:"name"`
	TextColor       string          `json:"text_color"`
	BackgroundColor string          `json:"background_color"`
	Tasks           []TasksResponse `json:"tasks"`
}
