package controller

import (
	"net/http"
	"strconv"
	"taskMaster/data/request"
	"taskMaster/data/response"
	"taskMaster/helper"
	"taskMaster/service"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

type TagsController struct {
	tagsService service.TagsService
}

func NewTagsController(service service.TagsService) *TagsController {
	return &TagsController{
		tagsService: service,
	}
}

// CreateTags		godoc
// @Summary			Create tags
// @Description		Save tags data in Db.
// @Param			tags body request.CreateTagsRequest true "Create tags"
// @Produce			application/json
// @Tags			tags
// @Success			200 {object} response.Response{}
// @Router			/tags [post]
func (controller *TagsController) Create(ctx *gin.Context) {
	log.Info().Msg("create tags")
	createTagsRequest := request.CreateTagsRequest{}
	err := ctx.ShouldBindJSON(&createTagsRequest)
	helper.ErrorPanic(err)

	id, err1 := controller.tagsService.Create(createTagsRequest)
	if err1 != nil {
		helper.HandleError(ctx, err1.Code, err1.Msg)
	} else {
		webResponse := response.Response{
			Code:   http.StatusOK,
			Status: "Ok",
			Data:   id,
		}
		ctx.JSON(http.StatusOK, webResponse)
	}
}

// UpdateTags		godoc
// @Summary			Update tags
// @Description		Update tags data.
// @Param			tagId path string true "update tags by id"
// @Param			tags body request.CreateTagsRequest true  "Update tags"
// @Tags			tags
// @Produce			application/json
// @Success			200 {object} response.Response{}
// @Router			/tags/{tagId} [patch]
func (controller *TagsController) Update(ctx *gin.Context) {
	log.Info().Msg("update tags")
	updateTagsRequest := request.UpdateTagsRequest{}
	err := ctx.ShouldBindJSON(&updateTagsRequest)
	helper.ErrorPanic(err)

	tagId := ctx.Param("tagId")
	id, err := strconv.Atoi(tagId)
	helper.ErrorPanic(err)
	updateTagsRequest.ID = uint(id)

	controller.tagsService.Update(updateTagsRequest)

	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   nil,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

// DeleteTags		godoc
// @Summary			Delete tags
// @Description		Remove tags data by id.
// @Produce			application/json
// @Tags			tags
// @Success			200 {object} response.Response{}
// @Router			/tags/{tagID} [delete]
func (controller *TagsController) Delete(ctx *gin.Context) {
	log.Info().Msg("delete tags")
	tagId := ctx.Param("tagId")
	id, err := strconv.Atoi(tagId)
	helper.ErrorPanic(err)
	controller.tagsService.Delete(uint(id))

	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   nil,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

// FindByIdTags 		godoc
// @Summary				Get Single tags by id.
// @Param				tagId path string true "update tags by id"
// @Description			Return the tahs whoes tagId valu mathes id.
// @Produce				application/json
// @Tags				tags
// @Success				200 {object} response.Response{}
// @Router				/tags/{tagId} [get]
func (controller *TagsController) FindById(ctx *gin.Context) {
	log.Info().Msg("findbyid tags")
	tagId := ctx.Param("tagId")
	id, err := strconv.Atoi(tagId)
	helper.ErrorPanic(err)
	tagResponse := controller.tagsService.FindById(uint(id))
	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   tagResponse,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

// FindAllTags 		godoc
// @Summary			Get All tags.
// @Description		Return list of tags.
// @Tags			tags
// @success 		200 {object} response.Response{data=[]response.TagsResponse} "Success"
// @Router			/tags [get]
func (controller *TagsController) FindAll(ctx *gin.Context) {
	log.Info().Msg("findAll tags")
	offsetStr := ctx.DefaultQuery("offset", "0")
	limitStr := ctx.DefaultQuery("limit", "50")
	offset, err := strconv.Atoi(offsetStr)
	if err != nil {
		// 如果转换失败，返回 400 错误
		ctx.JSON(http.StatusBadRequest, response.Response{
			Code:   http.StatusBadRequest,
			Status: "Error",
			Data:   "Invalid page parameter",
		})
		return
	}
	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		// 如果转换失败，返回 400 错误
		ctx.JSON(http.StatusBadRequest, response.Response{
			Code:   http.StatusBadRequest,
			Status: "Error",
			Data:   "Invalid size parameter",
		})
		return
	}
	tagResponse := controller.tagsService.FindAll(offset, limit)
	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   tagResponse,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)

}

// AssociateTasksWithTag godoc
// @Summary      Associate Tasks with Tag
// @Description  Add tasks to a tag by providing a tag ID and a list of task IDs.
// @Tags         tags
// @Accept       json
// @Produce      json
// @Param        tagId   path      uint      true  "Tag ID"
// @Param        body    body      []uint    true  "Array of Task IDs"
// @Success      200     {object}  response.Response{}
// @Failure      400     {object}  response.Response{}
// @Failure      500     {object}  response.Response{}
// @Router       /tags/{tagId}/tasks [post]
func (controller *TagsController) AssociateTasksWithTag(ctx *gin.Context) {
	tagId := ctx.Param("tagId")

	id, err := strconv.Atoi(tagId)
	if err != nil {
		helper.HandleError(ctx, http.StatusBadRequest, "Invalid tag ID")
	}

	var taskIds []uint
	if err := ctx.ShouldBindJSON(&taskIds); err != nil {
		helper.HandleError(ctx, http.StatusBadRequest, "Invalid task IDs")
	}

	controller.tagsService.AssociateTasksWithTag(uint(id), taskIds)

	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   nil,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

// DetachTasksFromTag godoc
// @Summary      Detach Tasks from Tag
// @Description  Remove tasks from a tag by providing a tag ID and a list of task IDs.
// @Tags         tags
// @Accept       json
// @Produce      json
// @Param        tagId   path      uint      true  "Tag ID"
// @Param        body    body      []uint    true  "Array of Task IDs"
// @Success      200     {object}  response.Response{}
// @Failure      400     {object}  response.Response{}
// @Failure      500     {object}  response.Response{}
// @Router       /tags/{tagId}/tasks [delete]
func (controller *TagsController) DetachTasksFromTag(ctx *gin.Context) {
	tagId := ctx.Param("tagId")

	id, err := strconv.Atoi(tagId)
	if err != nil {
		helper.HandleError(ctx, http.StatusBadRequest, "Invalid tag ID")
	}

	var taskIds []uint
	if err := ctx.ShouldBindJSON(&taskIds); err != nil {
		helper.HandleError(ctx, http.StatusBadRequest, "Invalid task IDs")
	}

	controller.tagsService.DetachTasksFromTag(uint(id), taskIds)

	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   nil,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}
