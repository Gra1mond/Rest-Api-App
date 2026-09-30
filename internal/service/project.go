package service

import (
	"context"
	"errors"
	"rest_api/internal/model"
	"rest_api/internal/repository"
)

type ProjectService struct{
	repo *repository.ProjectRepository
}

func NewProjectService(repo *repository.ProjectRepository)*ProjectService{
	return &ProjectService{
		repo: repo,
	}
}

func (s *ProjectService)Create(ctx context.Context,
project model.Project)(model.Project,error){
	if project.Name==""{
		return model.Project{},errors.New("name is required")
	}
	return s.repo.Create(ctx,project)
}

func (s *ProjectService)GetByID(ctx context.Context,
project_id int)(model.Project,error){
	return s.repo.GetByID(ctx,project_id)
}

func (s *ProjectService)GetByUserID(ctx context.Context,
user_id int)(model.Project,error){
	return s.repo.GetByUserID(ctx,user_id)
}

func (s *ProjectService)Update(ctx context.Context,
project model.Project)(model.Project,error){
	return s.repo.Update(ctx,project)
}

func (s *ProjectService)Delete(ctx context.Context,
project_id int)error{
	return s.repo.Delete(ctx,project_id)
}