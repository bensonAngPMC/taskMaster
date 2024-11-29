package tags

import "taskMaster/model"

type TagsRepository interface {
	Save(tags model.Tags) uint
	Update(tags model.Tags)
	Delete(tagsId uint)
	FindById(tagsId uint) (tags model.Tags, err error)
	FindAll(offset int, limit int) []model.Tags
	AssociateTasksWithTag(tagId uint, taskIds []uint)
	DetachTasksFromTag(tagId uint, taskIds []uint)
	FindByNameTextColorBackground(name, textColor, backgroundColor string) (*model.Tags, error)
}
	