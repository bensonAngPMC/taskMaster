package response

type TasksResponse struct {
	ID          uint            `json:"id"`
	PlannedDateTime string         `json:"planned_date_time"`
	ActualDateTime  string         `json:"actual_date_time"`
	IsDone      bool           `json:"is_done"`
	TimeDiff    string         `json:"time_diff"`
	Tags        []TagsResponse `json:"tags"`
	Title       string         `json:"title"`
	Description string         `json:"description"`
}
