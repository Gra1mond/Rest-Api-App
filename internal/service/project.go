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
	userID int,
project model.Project)(model.Project,error){
	if project.Name==""{
		return model.Project{},errors.New("name is required")
	}
	project.UserID = userID
	return s.repo.Create(ctx,project)
}

func (s *ProjectService)GetByID(ctx context.Context,
	UserID int,
ProjectID int)(model.Project,error){
	project,err:=s.repo.GetByID(ctx,ProjectID)
	if err!=nil{
		return model.Project{},err
	}
	if project.UserID!=UserID{
		return model.Project{},errors.New("access denied")
	}
	return s.repo.GetByID(ctx,ProjectID)
}

func (s *ProjectService)GetByUserID(ctx context.Context,
UserId int)(model.Project,error){
	return s.repo.GetByUserID(ctx,UserId)
}

func (s *ProjectService) Update(
	ctx context.Context,
	userID int,
	project model.Project,
) (model.Project, error) {

	if project.Name == "" {
		return model.Project{}, errors.New("name is required")
	}

	oldProject, err := s.repo.GetByID(ctx, project.ID)
	if err != nil {
		return model.Project{}, err
	}

	if oldProject.UserID != userID {
		return model.Project{}, errors.New("access denied")
	}

	return s.repo.Update(ctx, project)
}

func (s *ProjectService) Delete(
	ctx context.Context,
	userID int,
	ProjectID int,
) error {

	project,err:=s.repo.GetByID(ctx,ProjectID)
	if err!=nil{
		return err
	}
	if project.UserID!=userID{
		return errors.New("access denied")
	}
	return s.repo.Delete(ctx,ProjectID)
}