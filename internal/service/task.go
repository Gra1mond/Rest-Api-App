package service

import (
	"errors"
	"rest_api/internal/model"
	"rest_api/internal/repository"
)

type TaskService struct{
	repo *repository.TaskRepository
}

func NewService(repo *repository.TaskRepository)*TaskService{
	return &TaskService{
		repo: repo,
	}
}

func (s *TaskService) GetByID(id int)(model.Task,error){
	return s.repo.GetByID(id)
}

func (s* TaskService) GetAll()([]model.Task,error){
	return s.repo.GetAll()
}

func (s* TaskService) Create(task model.Task)(model.Task,error){
	if task.Title==""{
		return model.Task{},errors.New("title is required")
	}
	return s.repo.Create(task)
}