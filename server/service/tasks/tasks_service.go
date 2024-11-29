package tasks

import (
	"taskMaster/data/request"
	"taskMaster/data/response"
)

type TasksService interface {
	Create(tasks request.CreateTasksRequest) uint
	Update(tasks map[string]any)
	Delete(tasksId uint)
	FindById(tasksId uint, populate *map[string][]string) response.TasksResponse
	FindAll(populate *map[string][]string, tagIds []int) []response.TasksResponse
	AssociateTagsWithTask(taskId uint, tagIds []uint)
	DetachTagsFromTask(taskId uint, tagIds []uint)
}
