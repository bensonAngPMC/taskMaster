package tasks

import (
	"errors"
	"fmt"

	// "fmt"
	"taskMaster/helper"
	"taskMaster/model"
	"taskMaster/util"

	"gorm.io/gorm"
)

type TasksRepositoryImpl struct {
	Db *gorm.DB
}

// Delete implements TasksRepository
func (t *TasksRepositoryImpl) Delete(tasksId uint) {
	var tasks model.Tasks
	result := t.Db.Where("id = ?", tasksId).Delete(&tasks)
	helper.ErrorPanic(result.Error)
}

// FindAll implements TasksRepository
func (t *TasksRepositoryImpl) FindAll(populate *map[string][]string, tagIds []int) []model.Tasks {
	var tasks []model.Tasks
	db := t.Db
	if len(tagIds) > 0 {
		db = db.Joins("LEFT JOIN task_tags ON task_tags.tasks_id = tasks.id").
			Joins("LEFT JOIN tags ON task_tags.tags_id = tags.id").
			Where("tags.id IN (?)", tagIds)
	}
	if populate != nil {
		for key, fields := range *populate {
			tempKey := util.CapitalizeFirstLetter(key)
			if len(fields) > 0 {
				db = db.Preload(tempKey, func(db *gorm.DB) *gorm.DB {
					return db.Select(fields)
				})
			} else {
				db = db.Preload(tempKey)
			}
		}
	}

	result := db.Distinct("tasks.id, tasks.*").Order("is_done ASC, planned_date_time ASC").Find(&tasks)
	helper.ErrorPanic(result.Error)
	return tasks
}

// FindById implements TasksRepository
func (t *TasksRepositoryImpl) FindById(tasksId uint, populate *map[string][]string) (tasks model.Tasks, err error) {
	var task model.Tasks
	db := t.Db
	if populate != nil {
		for key, fields := range *populate {
			tempKey := util.CapitalizeFirstLetter(key)
			if len(fields) > 0 {
				db = db.Preload(tempKey, func(db *gorm.DB) *gorm.DB {
					return db.Select(fields)
				})
			} else {
				db = db.Preload(tempKey)
			}
			fmt.Println("tempKeytempKey", tempKey)
		}
	}
	result := db.Find(&task, tasksId)
	if result != nil {
		return task, nil
	} else {
		return task, errors.New("task is not found")
	}
}

// Save implements TasksRepository
func (t *TasksRepositoryImpl) Save(tasks model.Tasks) uint {
	result := t.Db.Create(&tasks)
	if result.Error != nil {
		helper.ErrorPanic(result.Error)
		return 0
	}
	return tasks.ID
}

// Update implements TasksRepository
func (t *TasksRepositoryImpl) Update(tasks model.Tasks) {
	result := t.Db.Model(&tasks).Select("*").Updates(tasks)
	helper.ErrorPanic(result.Error)
}

// AddTagsToTask
func (t *TasksRepositoryImpl) AssociateTagsWithTask(taskId uint, tagIds []uint) {
	var task model.Tasks
	var tags []model.Tags
	if err1 := t.Db.First(&task, taskId).Error; err1 != nil {
		helper.ErrorPanic(err1)
	}
	if err2 := t.Db.Find(&tags, tagIds).Error; err2 != nil {
		helper.ErrorPanic(err2)
	}
	if err3 := t.Db.Model(&task).Association("Tags").Append(tags); err3 != nil {
		helper.ErrorPanic(err3)
	}
}

// DetachTagsFromTask
func (t *TasksRepositoryImpl) DetachTagsFromTask(taskId uint, tagIds []uint) {
	var task model.Tasks
	var tags []model.Tags
	if err1 := t.Db.First(&task, taskId).Error; err1 != nil {
		helper.ErrorPanic(err1)
	}
	if err2 := t.Db.Find(&tags, tagIds).Error; err2 != nil {
		helper.ErrorPanic(err2)
	}
	if err3 := t.Db.Model(&task).Association("Tags").Delete(tags); err3 != nil {
		helper.ErrorPanic(err3)
	}
}
