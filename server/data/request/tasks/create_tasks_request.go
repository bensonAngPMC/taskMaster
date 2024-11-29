package tasks

type CreateTasksRequest struct {
	PlannedDateTime string `validate:"required,min=8,max=30" json:"planned_date_time"`
	ActualDateTime  string `validate:"omitempty,min=8,max=30" json:"actual_date_time"`
	IsDone          bool   `gorm:"type:boolean" json:"is_done"`
	TimeDiff        string `validate:"omitempty,min=1,max=50" json:"time_diff"`
	Title           string `validate:"required,min=1,max=100" json:"title"`
	Description     string `validate:"omitempty,min=1,max=500" json:"description"`
}
