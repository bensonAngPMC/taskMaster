package controller

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"taskMaster/data/request"
	"taskMaster/data/response"
	"taskMaster/helper"
	"taskMaster/service"
	"taskMaster/util"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

type TasksController struct {
	tasksService service.TasksService
}

func NewTasksController(service service.TasksService) *TasksController {
	return &TasksController{
		tasksService: service,
	}
}

// CreateTasks		godoc
// @Summary			Create tasks
// @Description		Save tasks data in Db.
// @Param			tasks body request.CreateTasksRequest true "Create tasks"
// @Produce			application/json
// @Tags			tasks
// @Success			200 {object} response.Response{}
// @Router			/tasks [post]
func (controller *TasksController) Create(ctx *gin.Context) {
	log.Info().Msg("create tasks")
	createTasksRequest := request.CreateTasksRequest{}
	err := ctx.ShouldBindJSON(&createTasksRequest)
	helper.ErrorPanic(err)

	id := controller.tasksService.Create(createTasksRequest)
	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   id,
	}
	ctx.JSON(http.StatusOK, webResponse)
}

// UpdateTasks		godoc
// @Summary			Update tasks
// @Description		Update tasks data.
// @Param			taskId path string true "update tasks by id"
// @Param			tasks body request.CreateTasksRequest true  "Update tasks"
// @Tags			tasks
// @Produce			application/json
// @Success			200 {object} response.Response{}
// @Router			/tasks/{taskId} [patch]
func (controller *TasksController) Update(ctx *gin.Context) {
	log.Info().Msg("update tasks")
	// updateTasksRequest := request.UpdateTasksRequest{}
	var updateTasksRequest map[string]any
	// var updateTasksRequest map[string]interface{}
	err := ctx.ShouldBindJSON(&updateTasksRequest)
	fmt.Println(updateTasksRequest)
	helper.ErrorPanic(err)

	taskId := ctx.Param("taskId")
	fmt.Println(taskId)
	id, err := strconv.Atoi(taskId)
	helper.ErrorPanic(err)
	updateTasksRequest["ID"] = uint(id)

	controller.tasksService.Update(updateTasksRequest)

	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   nil,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

// DeleteTasks		godoc
// @Summary			Delete tasks
// @Description		Remove tasks data by id.
// @Produce			application/json
// @Tags			tasks
// @Success			200 {object} response.Response{}
// @Router			/tasks/{taskID} [delete]
func (controller *TasksController) Delete(ctx *gin.Context) {
	log.Info().Msg("delete tasks")
	taskId := ctx.Param("taskId")
	id, err := strconv.Atoi(taskId)
	helper.ErrorPanic(err)
	controller.tasksService.Delete(uint(id))

	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   nil,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

// FindByIdTasks 		godoc
// @Summary				Get Single tasks by id.
// @Param				taskId path string true "update tasks by id"
// @Description			Return the tahs whoes taskId valu mathes id.
// @Produce				application/json
// @Tags				tasks
// @Success				200 {object} response.Response{}
// @Router				/tasks/{taskId} [get]
func (controller *TasksController) FindById(ctx *gin.Context) {
	log.Info().Msg("findbyid tasks")

	taskId := ctx.Param("taskId")
	populate := ctx.Request.URL.Query()["populate"] // 获取多个 `populate` 参数
	populateMap := util.ParsePopulate(populate)
	id, err := strconv.Atoi(taskId)
	helper.ErrorPanic(err)
	taskResponse := controller.tasksService.FindById(uint(id), &populateMap)
	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   taskResponse,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

// FindAllTasks 	godoc
// @Summary			Get All tasks.
// @Description		Return list of tasks.
// @Tags			tasks
// @Success			200 {obejct} response.Response{}
// @Router			/tasks [get]
func (controller *TasksController) FindAll(ctx *gin.Context) {
	log.Info().Msg("findAll tasks")
	populate := ctx.Request.URL.Query()["populate"] // 获取多个 `populate` 参数
	populateMap := util.ParsePopulate(populate)

	tagsParam := ctx.DefaultQuery("tags.id_in", "") // 获取 tags.id_in 参数，默认值为空字符串
	var tagIds []int
	if tagsParam != "" {
		// 将逗号分隔的 tags.id_in 字符串转换为整数数组
		tagsStr := strings.Split(tagsParam, ",")
		for _, id := range tagsStr {
			tagId, err := strconv.Atoi(id)
			if err == nil {
				tagIds = append(tagIds, tagId)
			}
		}
	}
	taskResponse := controller.tasksService.FindAll(&populateMap, tagIds)
	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   taskResponse,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

// AssociateTagsWithTask godoc
// @Summary      Associate Tags with Task
// @Description  Add tags to a task by providing a task ID and a list of tag IDs.
// @Tags         tasks
// @Accept       json
// @Produce      json
// @Param        taskId  path      uint      true  "Task ID"
// @Param        body    body      []uint    true  "Array of Tag IDs"
// @Success      200     {object}  response.Response{}
// @Failure      400     {object}  response.Response{}
// @Failure      500     {object}  response.Response{}
// @Router       /tasks/{taskId}/tags [post]
func (controller *TasksController) AssociateTagsWithTask(ctx *gin.Context) {
	taskId := ctx.Param("taskId")

	id, err := strconv.Atoi(taskId)
	if err != nil {
		helper.HandleError(ctx, http.StatusBadRequest, "Invalid task ID")
	}

	var tagIds []uint
	if err := ctx.ShouldBindJSON(&tagIds); err != nil {
		helper.HandleError(ctx, http.StatusBadRequest, "Invalid tag IDs")
	}

	controller.tasksService.AssociateTagsWithTask(uint(id), tagIds)

	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   nil,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

// DetachTagsFromTask godoc
// @Summary      Detach Tags from Task
// @Description  Remove tags from a task by providing a task ID and a list of tag IDs.
// @Tags         tasks
// @Accept       json
// @Produce      json
// @Param        taskId  path      uint      true  "Task ID"
// @Param        body    body      []uint    true  "Array of Tag IDs"
// @Success      200     {object}  response.Response{}
// @Failure      400     {object}  response.Response{}
// @Failure      500     {object}  response.Response{}
// @Router       /tasks/{taskId}/tags [delete]
func (controller *TasksController) DetachTagsFromTask(ctx *gin.Context) {
	taskId := ctx.Param("taskId")

	id, err := strconv.Atoi(taskId)
	if err != nil {
		helper.HandleError(ctx, http.StatusBadRequest, "Invalid task ID")
	}

	var tagIds []uint
	if err := ctx.ShouldBindJSON(&tagIds); err != nil {
		helper.HandleError(ctx, http.StatusBadRequest, "Invalid tag IDs")
	}

	controller.tasksService.DetachTagsFromTask(uint(id), tagIds)

	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   nil,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}
