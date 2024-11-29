package service

import (
	"taskMaster/repository"
	"taskMaster/service/tags"
	"taskMaster/service/tasks"

	"github.com/go-playground/validator/v10"
)

type (
	TagsService = tags.TagsService
	TasksService = tasks.TasksService
)

func NewTagsServiceImpl(tagRepository repository.TagsRepository, validate *validator.Validate) TagsService {
	return &tags.TagsServiceImpl{
		TagsRepository: tagRepository,
		Validate:       validate,
	}
}
func NewTasksServiceImpl(taskRepository repository.TasksRepository, validate *validator.Validate) TasksService {
	return &tasks.TasksServiceImpl{
		TasksRepository: taskRepository,
		Validate:        validate,
	}
}
