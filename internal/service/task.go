package service

import (
	"context"
	"errors"
	"rest_api/internal/model"
	"rest_api/internal/repository"
)

type TaskService struct{
	taskRepo    *repository.TaskRepository
	projectRepo *repository.ProjectRepository
}

func NewTaskService(
	taskRepo *repository.TaskRepository,
	projectRepo *repository.ProjectRepository,
) *TaskService {
	return &TaskService{
		taskRepo:    taskRepo,
		projectRepo: projectRepo,
	}
}

func (s *TaskService)GetByID(ctx context.Context,
	userID int,
task_id int)(model.Task,error){
	
	task,err:=s.taskRepo.GetByID(ctx,task_id)
	if err!=nil{
		return model.Task{},err
	}
	project,err:=s.projectRepo.GetByID(ctx,task.ProjectID)
	if err!=nil{
		return model.Task{},err
	}
	if project.UserID!=userID{
		return model.Task{},errors.New("access denied")
	}
	return task,nil
}

func (s *TaskService)Create(ctx context.Context,
	userID int,
task model.Task,
projectID int,)(model.Task,error){
	if task.Title==""{
		return model.Task{},errors.New("title is required")
	}
	project,err:=s.projectRepo.GetByID(ctx,projectID)
	if err!=nil{
		return model.Task{},err
	}

	if project.UserID!=userID{
		return model.Task{},errors.New("access denied")
	}

	task.ProjectID=projectID
	return s.taskRepo.Create(ctx,task)
}

func (s *TaskService)GetByProjectID(ctx context.Context,
	UserID int,
project_id int)([]model.Task,error){
	project,err:=s.projectRepo.GetByID(ctx,project_id)
	if err!=nil{
		return []model.Task{},err
	}
	if project.UserID!=UserID{
		return []model.Task{},errors.New("access denied")
	}
	return s.taskRepo.GetByProjectID(ctx,project_id)
}

func (s *TaskService)Update(ctx context.Context,
	UserID int,
task model.Task)(model.Task,error){
	if task.Title==""{
		return model.Task{},errors.New("title is required")
	}
	olTask,err:=s.taskRepo.GetByID(ctx,task.ID)
	if err!=nil{
		return model.Task{},err
	}

	project,err:=s.projectRepo.GetByID(ctx,olTask.ProjectID)
	if err!=nil{
		return model.Task{},err
	}
	if project.UserID!=UserID{
		return model.Task{},errors.New("access denied")
	}
	task.ProjectID=olTask.ProjectID
	return s.taskRepo.Update(ctx,task)
}

func (s *TaskService)Delete(ctx context.Context,
	UserID int,
task_id int)error{
	task,err:=s.taskRepo.GetByID(ctx,task_id)
	if err!=nil{
		return err
	}
	project,err:=s.projectRepo.GetByID(ctx,task.ProjectID)
	if err!=nil{
		return err
	}
	if project.UserID!=UserID{
		return errors.New("access denied")
	}
	
	return s.taskRepo.Delete(ctx,task_id)
}