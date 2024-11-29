package tags

type UpdateTagsRequest struct {
	ID              uint    `validate:"required"`
	Name            string `validate:"omitempty,min=1,max=255" json:"name"`
	TextColor       string `validate:"omitempty,min=7,max=7" json:"text_color"`
	BackgroundColor string `validate:"omitempty,min=7,max=7" json:"background_color"`
	TaskIDs         []uint  `json:"task_ids"`
}
