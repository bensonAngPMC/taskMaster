package repository

import (
	"taskMaster/repository/tags"
	"taskMaster/repository/tasks"

	"gorm.io/gorm"
)

type (
	TagsRepository  = tags.TagsRepository
	TasksRepository = tasks.TasksRepository
)

func NewTagsRepositoryImpl(Db *gorm.DB) TagsRepository {
	return &tags.TagsRepositoryImpl{Db: Db}
}

func NewTasksRepositoryImpl(Db *gorm.DB) TasksRepository {
	return &tasks.TasksRepositoryImpl{Db: Db}
}
