package main

import (
	"net/http"
	"taskMaster/config"
	"taskMaster/controller"
	_ "taskMaster/docs"
	"taskMaster/helper"
	"taskMaster/model"
	"taskMaster/repository"
	"taskMaster/router"
	"taskMaster/service"

	// "time"

	// "github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/rs/zerolog/log"
	// "github.com/gin-contrib/cors"
)

// @title 	TaskMaster Service API
// @version	1.0
// @description A TaskMaster service API in Go using Gin framework

// @host 	localhost:8888
// @BasePath /api
func main() {
	log.Info().Msg("Started Server!")
	// Database
	db := config.DatabaseConnection()
	validate := validator.New()
	db.AutoMigrate(&model.Tags{}, &model.Tasks{})
	// Repository
	tagsRepository := repository.NewTagsRepositoryImpl(db)
	tasksRepository := repository.NewTasksRepositoryImpl(db)
	// Service
	tagsService := service.NewTagsServiceImpl(tagsRepository, validate)
	tasksService := service.NewTasksServiceImpl(tasksRepository, validate)
	// Controller
	tagsController := controller.NewTagsController(tagsService)
	tasksController := controller.NewTasksController(tasksService)
	// Router
	routes := router.NewRouter(tagsController, tasksController)
	server := &http.Server{
		Addr:    ":8888",
		Handler: routes,
	}
	err := server.ListenAndServe()
	helper.ErrorPanic(err)
}
