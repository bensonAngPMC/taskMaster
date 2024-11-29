package tasks

import (
	"taskMaster/data/request"
	"taskMaster/data/response"
	"taskMaster/helper"
	"taskMaster/model"
	"taskMaster/repository"

	"github.com/go-playground/validator/v10"
)

type TasksServiceImpl struct {
	TasksRepository repository.TasksRepository
	Validate        *validator.Validate
}

// Create implements TasksService
func (t *TasksServiceImpl) Create(tasks request.CreateTasksRequest) uint {
	var tempId uint
	err := t.Validate.Struct(tasks)
	helper.ErrorPanic(err)
	taskModel := model.Tasks{
		PlannedDateTime: tasks.PlannedDateTime,
		ActualDateTime:  tasks.ActualDateTime,
		IsDone:          tasks.IsDone,
		TimeDiff:        tasks.TimeDiff,
		Title:           tasks.Title,
		Description:     tasks.Description,
	}
	tempId = t.TasksRepository.Save(taskModel)
	return tempId
}

// Delete implements TasksService
func (t *TasksServiceImpl) Delete(tasksId uint) {
	t.TasksRepository.Delete(tasksId)
}

// FindAll implements TasksService
func (t *TasksServiceImpl) FindAll(populate *map[string][]string, tagIds []int) []response.TasksResponse {
	result := t.TasksRepository.FindAll(populate, tagIds)

	var tasks []response.TasksResponse
	for _, value := range result {
		task := response.TasksResponse{
			ID:              value.ID,
			PlannedDateTime: value.PlannedDateTime,
			ActualDateTime:  value.ActualDateTime,
			IsDone:          value.IsDone,
			TimeDiff:        value.TimeDiff,
			Title:           value.Title,
			Description:     value.Description,
			Tags:            []response.TagsResponse{},
		}

		for _, tagsItem := range value.Tags {
			tag := response.TagsResponse{
				ID:              tagsItem.ID,
				Name:            tagsItem.Name,
				TextColor:       tagsItem.TextColor,
				BackgroundColor: tagsItem.BackgroundColor,
			}

			task.Tags = append(task.Tags, tag)
		}
		tasks = append(tasks, task)
	}

	return tasks
}

// FindById implements TasksService
func (t *TasksServiceImpl) FindById(tasksId uint, populate *map[string][]string) response.TasksResponse {
	taskData, err := t.TasksRepository.FindById(tasksId, populate)
	helper.ErrorPanic(err)

	taskResponse := response.TasksResponse{
		ID:              taskData.ID,
		PlannedDateTime: taskData.PlannedDateTime,
		ActualDateTime:  taskData.ActualDateTime,
		IsDone:          taskData.IsDone,
		TimeDiff:        taskData.TimeDiff,
		Title:           taskData.Title,
		Description:     taskData.Description,
		Tags:            []response.TagsResponse{},
	}

	for _, tagsItem := range taskData.Tags {
		tag := response.TagsResponse{
			ID:              tagsItem.ID,
			Name:            tagsItem.Name,
			TextColor:       tagsItem.TextColor,
			BackgroundColor: tagsItem.BackgroundColor,
		}

		taskResponse.Tags = append(taskResponse.Tags, tag)
	}
	return taskResponse
}

// Update implements TasksService
func (t *TasksServiceImpl) Update(tasks map[string]any) {
	// Fetch the task from the repository by ID
	// var taskID uint
	// if id, ok := tasks["ID"].(string); ok {
	// 	tasksIDInt, _ := strconv.Atoi(id)
	// 	taskID = uint(tasksIDInt)
	// } else if id, ok := tasks["ID"].(uint); ok {
	// 	taskID = id
	// } else {
	// 	helper.ErrorPanic(fmt.Errorf("ID field is missing or invalid"))
	// }

	// taskData, err := t.TasksRepository.FindById(taskID)
	// helper.ErrorPanic(err)

	// // Assign other fields if present
	// // taskData.ID = taskID
	// if plannedDateTime, ok := tasks["planned_date_time"].(string); ok {
	// 	taskData.PlannedDateTime = plannedDateTime
	// }

	// if actualDateTime, ok := tasks["actual_date_time"].(string); ok {
	// 	taskData.ActualDateTime = actualDateTime
	// }

	// if isDone, ok := tasks["is_done"].(bool); ok {
	// 	taskData.IsDone = isDone
	// }

	// if timeDiff, ok := tasks["time_diff"].(string); ok {
	// 	taskData.TimeDiff = timeDiff
	// }

	// if title, ok := tasks["title"].(string); ok {
	// 	taskData.Title = title
	// }

	// if description, ok := tasks["description"].(string); ok {
	// 	taskData.Description = description
	// }
	// fmt.Println("taskData:", taskData)

	// // Save the updated task data
	// t.TasksRepository.Update(taskData)
}

func (t *TasksServiceImpl) AssociateTagsWithTask(taskId uint, tagIds []uint) {
	t.TasksRepository.AssociateTagsWithTask(taskId, tagIds)
}
func (t *TasksServiceImpl) DetachTagsFromTask(taskId uint, tagIds []uint) {
	t.TasksRepository.DetachTagsFromTask(taskId, tagIds)
}
