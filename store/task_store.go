package store

import "stability-test-task-api/models"

var Tasks = []models.Task{
	{ID: 1, Title: "Learn Go", Done: false},
	{ID: 2, Title: "Build API", Done: false},
}

func GetAllTasks() []models.Task {
	return Tasks
}

func GetTaskByID(id int) *models.Task {
	for i := range Tasks {
		if Tasks[i].ID == id {
			return &Tasks[i]
		}
	}
	return nil
}

func AddTask(task models.Task) {
	task.ID = Tasks[len(Tasks)-1].ID + 1
	Tasks = append(Tasks, task)
}

func DeleteTask(id int) bool {
	for i, t := range Tasks {
		if t.ID == id {
			Tasks = append(Tasks[:i], Tasks[i+1:]...)
			return true
		}
	}
	return false
}
