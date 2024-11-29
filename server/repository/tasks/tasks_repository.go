package tasks

import "taskMaster/model"

type TasksRepository interface {
	Save(tasks model.Tasks) uint
	Update(tasks model.Tasks)
	Delete(tasksId uint)
	FindById(tasksId uint, populate *map[string][]string) (tasks model.Tasks, err error)
	FindAll(populate *map[string][]string, tagIds []int) []model.Tasks
	AssociateTagsWithTask(taskId uint, tagIds []uint)
	DetachTagsFromTask(taskId uint, tagIds []uint)
}
