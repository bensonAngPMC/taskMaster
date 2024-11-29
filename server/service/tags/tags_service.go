package tags

import (
	"taskMaster/data/request"
	"taskMaster/data/response"
	"taskMaster/helper"
)

type TagsService interface {
	Create(tags request.CreateTagsRequest) (id uint ,ErrorResponse *helper.ErrorResponse)
	Update(tags request.UpdateTagsRequest)
	Delete(tagsId uint)
	FindById(tagsId uint) response.TagsResponse
	FindAll(offset int, limit int) []response.TagsResponse
	AssociateTasksWithTag(tagId uint, taskIds []uint)
	DetachTasksFromTag(tagId uint, taskIds []uint)
}
