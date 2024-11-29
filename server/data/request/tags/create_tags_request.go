package tags

type CreateTagsRequest struct {
	Name            string                     `validate:"required,min=1,max=255" json:"name"`
	TextColor       string                     `validate:"required,min=7,max=7" json:"text_color"`
	BackgroundColor string                     `validate:"required,min=7,max=7" json:"background_color"`
}
