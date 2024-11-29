package request

import (
	"taskMaster/data/request/tags"
	"taskMaster/data/request/tasks"
)

type (
	// Tags
	CreateTagsRequest  = tags.CreateTagsRequest
	UpdateTagsRequest  = tags.UpdateTagsRequest
	// Tasks
	CreateTasksRequest = tasks.CreateTasksRequest
	UpdateTasksRequest = tasks.UpdateTasksRequest
	UpdateTasksDBRequest = tasks.UpdateTasksDBRequest
)
