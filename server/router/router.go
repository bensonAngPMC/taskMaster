package router

import (
	"net/http"
	"taskMaster/controller"
	"taskMaster/middleware"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewRouter(tagsController *controller.TagsController, tasksController *controller.TasksController) *gin.Engine {
	router := gin.Default()
	router.Use(middleware.Cors())
	// add swagger
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.GET("", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "welcome home")
	})

	baseRouter := router.Group("/api")
	tagsRouter := baseRouter.Group("/tags")
	tagsRouter.GET("", tagsController.FindAll)
	tagsRouter.GET("/:tagId", tagsController.FindById)
	tagsRouter.POST("", tagsController.Create)
	tagsRouter.POST("/:tagId/tags", tagsController.AssociateTasksWithTag)
	tagsRouter.PATCH("/:tagId", tagsController.Update)
	tagsRouter.DELETE("/:tagId", tagsController.Delete)
	tagsRouter.DELETE("/:tagId/tags", tagsController.DetachTasksFromTag)

	tasksRouter := baseRouter.Group("/tasks")
	tasksRouter.GET("", tasksController.FindAll)
	tasksRouter.GET("/:taskId", tasksController.FindById)
	tasksRouter.POST("", tasksController.Create)
	tasksRouter.POST("/:taskId/tags", tasksController.AssociateTagsWithTask)
	tasksRouter.PATCH("/:taskId", tasksController.Update)
	tasksRouter.DELETE("/:taskId", tasksController.Delete)
	tasksRouter.DELETE("/:taskId/tags", tasksController.DetachTagsFromTask)
	return router
}
