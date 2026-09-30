package repository

import (
	"context"
	"rest_api/internal/model"

	"github.com/jackc/pgx/v5"
)

type ProjectRepository struct{
	db *pgx.Conn
}

func (r *ProjectRepository)Create(ctx context.Context,
project model.Project)(model.Project,error){
	err:=r.db.QueryRow(ctx,
	`INSERT INTO project(user_id,name)
	RETURNING id`,
project.UserID,
project.Name).Scan(&project.ID)
return project,err
}

func (r *ProjectRepository) GetByID(ctx context.Context,
id int)(model.Project,error){
	var project model.Project
	err:=r.db.QueryRow(ctx,
	`SELECT user_id,name
	FROM projects
	WHERE id = $1`,
id).Scan(&project.UserID,&project.Name)
return project,err
}

func (r *ProjectRepository) GetByUserID(ctx context.Context,
user_id int)(model.Project,error){
	var project model.Project
	err:=r.db.QueryRow(ctx,
	`SELECT id,name
	FROM projects
	WHERE user_id = $1`,
user_id).Scan(&project.ID,&project.Name)
return project,err
}

func (r *ProjectRepository) Update(ctx context.Context,
project model.Project)(model.Project,error){
	err:=r.db.QueryRow(ctx,
	`UPDATE projects
	SET name = $1
	WHERE id = $2
	RETURNING id,user_id,name`,project.Name,project.ID).Scan(
		&project.ID,
		&project.UserID,
		&project.Name,
	)
	if err!=nil{
		return model.Project{},err
	}
	return project,nil
}