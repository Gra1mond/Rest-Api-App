package service

import (
	"context"
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

func (s *TaskService)GetByID(ctx context.Context,
task_id int)(model.Task,error){
	return s.repo.GetByID(ctx,task_id)
}

func (s *TaskService)Create(ctx context.Context,
task model.Task,
projectID int,)(model.Task,error){
	if task.Title==""{
		return model.Task{},errors.New("title is required")
	}
	task.ProjectID=projectID
	return s.repo.Create(ctx,task)
}

func (s *TaskService)GetByProjectID(ctx context.Context,
project_id int)([]model.Task,error){
	return s.repo.GetByProjectID(ctx,project_id)
}

func (s *TaskService)Update(ctx context.Context,
task model.Task)(model.Task,error){
	if task.Title==""{
		return model.Task{},errors.New("title is required")
	}
	return s.repo.Update(ctx,task)
}

func (s *TaskService)Delete(ctx context.Context,
task_id int)error{
	return s.repo.Delete(ctx,task_id)
}