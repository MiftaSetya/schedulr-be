package repository

import (
	"schedulr/internal/models"
	"time"

	"gorm.io/gorm"
)

type TaskRepository struct {
	DB *gorm.DB
}

func (r *TaskRepository) CreateTask(task *models.Task) error {
	return r.DB.Create(task).Error
}

func (r *TaskRepository) GetTaskById(id uint) (*models.Task, error) {
	task := &models.Task{}
	err := r.DB.First(task, id).Error
	return task, err
}

func (r *TaskRepository) GetTasksByListId(listID uint) ([]models.Task, error) {
	var tasks []models.Task
	err := r.DB.Where("list_id = ?", listID).Order("position asc").Find(&tasks).Error
	return tasks, err
}

func (r *TaskRepository) GetDueTasks(before time.Time) ([]models.Task, error) {
	var tasks []models.Task
	err := r.DB.Where("due_time <= ? AND completed = ?", before, false).Find(&tasks).Error
	return tasks, err
}

func (r *TaskRepository) ToggleCompleted(taskID uint, completed bool) error {
	return r.DB.Model(&models.Task{}).
		Where("id = ?", taskID).
		Update("completed", completed).Error
}


func (r *TaskRepository) UpdateTask(task *models.Task) error {
	return r.DB.Save(task).Error
}

func (r *TaskRepository) UpdateTaskPosition(taskID uint, newPos int) error {
	return r.DB.Model(&models.Task{}).
		Where("id = ?", taskID).
		Update("position", newPos).Error
}

func (r *TaskRepository) DeleteTask(id uint) error {
	return r.DB.Delete(&models.Task{}, id).Error
}
