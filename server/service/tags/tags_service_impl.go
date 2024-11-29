package tags

import (
	"taskMaster/data/request"
	"taskMaster/data/response"
	"taskMaster/helper"
	"taskMaster/model"
	"taskMaster/repository"

	"github.com/go-playground/validator/v10"
)

type TagsServiceImpl struct {
	TagsRepository repository.TagsRepository
	Validate       *validator.Validate
}

// Create implements TagsService
func (t *TagsServiceImpl) Create(tags request.CreateTagsRequest) (uint, *helper.ErrorResponse) {
	// 初始化返回的 ID
	var tempId uint

	// 验证输入数据
	err := t.Validate.Struct(tags)
	if err != nil {
		return tempId, &helper.ErrorResponse{
			Code: 400,
			Msg:  "Invalid input data",
		}
	}

	// 创建标签模型
	tagModel := model.Tags{
		Name:            tags.Name,
		TextColor:       tags.TextColor,
		BackgroundColor: tags.BackgroundColor,
	}

	// 查找现有标签
	existingTag, err := t.TagsRepository.FindByNameTextColorBackground(tagModel.Name, tagModel.TextColor, tagModel.BackgroundColor)
	if err != nil {
		if existingTag != nil {
			return tempId, &helper.ErrorResponse{
				Code: 409,
				Msg:  "Tag already exists",
			}
		}
	}

	// 如果标签已存在，返回错误

	// 保存新标签
	tempId = t.TagsRepository.Save(tagModel)
	return tempId, nil
}

// Delete implements TagsService
func (t *TagsServiceImpl) Delete(tagsId uint) {
	t.TagsRepository.Delete(tagsId)
}

// FindAll implements TagsService
func (t *TagsServiceImpl) FindAll(offset int, size int) []response.TagsResponse {
	// 获取分页后的数据
	result := t.TagsRepository.FindAll(offset, size)

	// 构建返回的标签列表
	var tags []response.TagsResponse
	for _, value := range result {
		tag := response.TagsResponse{
			ID:              value.ID,
			Name:            value.Name,
			TextColor:       value.TextColor,
			BackgroundColor: value.BackgroundColor,
			Tasks:           []response.TasksResponse{},
		}

		for _, tasksItem := range value.Tasks {
			task := response.TasksResponse{
				ID:              tasksItem.ID,
				PlannedDateTime: tasksItem.PlannedDateTime,
				ActualDateTime:  tasksItem.ActualDateTime,
				IsDone:          tasksItem.IsDone,
				TimeDiff:        tasksItem.TimeDiff,
				Title:           tasksItem.Title,
				Description:     tasksItem.Description,
			}

			tag.Tasks = append(tag.Tasks, task)
		}
		tags = append(tags, tag)
	}

	return tags
}

// FindById implements TagsService
func (t *TagsServiceImpl) FindById(tagsId uint) response.TagsResponse {
	tagData, err := t.TagsRepository.FindById(tagsId)
	helper.ErrorPanic(err)

	tagResponse := response.TagsResponse{
		ID:              tagData.ID,
		Name:            tagData.Name,
		TextColor:       tagData.TextColor,
		BackgroundColor: tagData.BackgroundColor,
		Tasks:           []response.TasksResponse{},
	}
	for _, tasksItem := range tagData.Tasks {
		task := response.TasksResponse{
			ID:              tasksItem.ID,
			PlannedDateTime: tasksItem.PlannedDateTime,
			ActualDateTime:  tasksItem.ActualDateTime,
			IsDone:          tasksItem.IsDone,
			TimeDiff:        tasksItem.TimeDiff,
			Title:           tasksItem.Title,
			Description:     tasksItem.Description,
		}

		tagResponse.Tasks = append(tagResponse.Tasks, task)
	}
	return tagResponse
}

// Update implements TagsService
func (t *TagsServiceImpl) Update(tags request.UpdateTagsRequest) {
	tagData, err := t.TagsRepository.FindById(tags.ID)
	helper.ErrorPanic(err)
	tagData.Name = tags.Name
	tagData.TextColor = tags.TextColor
	tagData.BackgroundColor = tags.BackgroundColor
	t.TagsRepository.Update(tagData)
}
func (t *TagsServiceImpl) AssociateTasksWithTag(tagId uint, taskIds []uint) {
	t.TagsRepository.AssociateTasksWithTag(tagId, taskIds)
}
func (t *TagsServiceImpl) DetachTasksFromTag(tagId uint, taskIds []uint) {
	t.TagsRepository.DetachTasksFromTag(tagId, taskIds)
}
