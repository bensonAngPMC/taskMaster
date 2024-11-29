package tasks

type UpdateTasksRequest struct {
	ID              uint   `validate:"required"`
	PlannedDateTime string `validate:"omitempty,min=8,max=30" json:"planned_date_time"`
	ActualDateTime  string `validate:"omitempty,min=8,max=30" json:"actual_date_time"`
	IsDone          bool   `gorm:"type:boolean" json:"is_done"`
	TimeDiff        string `validate:"omitempty,min=1,max=50" json:"time_diff"`
	TagIDs          []uint `json:"tag_ids"`
	Title           string `validate:"omitempty,min=1,max=100" json:"title"`
	Description     string `validate:"omitempty,min=1,max=500" json:"description"`
}

