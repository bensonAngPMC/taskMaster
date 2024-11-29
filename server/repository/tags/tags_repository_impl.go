package tags

import (
	"errors"
	"taskMaster/data/request"
	"taskMaster/helper"
	"taskMaster/model"

	"gorm.io/gorm"
)

type TagsRepositoryImpl struct {
	Db *gorm.DB
}

// Delete implements TagsRepository
func (t *TagsRepositoryImpl) Delete(tagsId uint) {
	var tags model.Tags
	result := t.Db.Where("id = ?", tagsId).Delete(&tags)
	helper.ErrorPanic(result.Error)
}

// FindAll implements TagsRepository
func (t *TagsRepositoryImpl) FindAll(offset int, limit int) []model.Tags {
	var tags []model.Tags
	result := t.Db.Offset(offset).Limit(limit).Find(&tags)
	helper.ErrorPanic(result.Error)
	return tags
}

// FindById implements TagsRepository
func (t *TagsRepositoryImpl) FindById(tagsId uint) (tags model.Tags, err error) {
	var tag model.Tags
	result := t.Db.Find(&tag, tagsId)
	if result != nil {
		return tag, nil
	} else {
		return tag, errors.New("tag is not found")
	}
}

// FindByNameTextColorBackground implements TagsRepository
func (t *TagsRepositoryImpl) FindByNameTextColorBackground(name string, textColor string, backgroundColor string) (*model.Tags, error) {
	var tag model.Tags
	err := t.Db.Where("name = ? AND text_color = ? AND background_color = ?", name, textColor, backgroundColor).First(&tag).Error
	if err != nil {
		// if gorm.IsRecordNotFoundError(err) {
		//     return nil, nil
		// }
		return nil, err
	}
	return &tag, nil
}

// Save implements TagsRepository
func (t *TagsRepositoryImpl) Save(tags model.Tags) uint {
    result := t.Db.Create(&tags)
    if result.Error != nil {
        helper.ErrorPanic(result.Error)
        return 0
    }

    return tags.ID 
}

// Update implements TagsRepository
func (t *TagsRepositoryImpl) Update(tags model.Tags) {
	var updateTag = request.UpdateTagsRequest{
		ID:              tags.ID,
		Name:            tags.Name,
		TextColor:       tags.TextColor,
		BackgroundColor: tags.BackgroundColor,
	}
	result := t.Db.Model(&tags).Updates(updateTag)
	helper.ErrorPanic(result.Error)
}

// AddTasksToTag
func (t *TagsRepositoryImpl) AssociateTasksWithTag(tagId uint, taskIds []uint) {
	var tasks []model.Tasks
	var tag model.Tags
	if err1 := t.Db.First(&tag, tagId).Error; err1 != nil {
		helper.ErrorPanic(err1)
	}
	if err2 := t.Db.Find(&tasks, taskIds).Error; err2 != nil {
		helper.ErrorPanic(err2)
	}
	if err3 := t.Db.Model(&tag).Association("Tasks").Append(tasks); err3 != nil {
		helper.ErrorPanic(err3)
	}
}

// AddTasksToTag
func (t *TagsRepositoryImpl) DetachTasksFromTag(tagId uint, taskIds []uint) {
	var tasks []model.Tasks
	var tag model.Tags
	if err1 := t.Db.First(&tag, tagId).Error; err1 != nil {
		helper.ErrorPanic(err1)
	}
	if err2 := t.Db.Find(&tasks, taskIds).Error; err2 != nil {
		helper.ErrorPanic(err2)
	}
	if err3 := t.Db.Model(&tag).Association("Tasks").Delete(tasks); err3 != nil {
		helper.ErrorPanic(err3)
	}
}
