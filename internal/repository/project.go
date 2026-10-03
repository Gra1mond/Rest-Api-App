package repository

import (
	"context"
	"errors"
	"rest_api/internal/model"

	"github.com/jackc/pgx/v5"
)

type ProjectRepository struct {
	db *pgx.Conn
}

func NewProjectRepository(db *pgx.Conn) *ProjectRepository {
	return &ProjectRepository{
		db: db,
	}
}

func (r *ProjectRepository) Create(ctx context.Context,
	project model.Project) (model.Project, error) {
	err := r.db.QueryRow(ctx,
		`INSERT INTO projects(user_id,name)
	VALUES($1, $2)
	RETURNING id`,
		project.UserID,
		project.Name).Scan(&project.ID)
	return project, err
}

func (r *ProjectRepository) GetByID(ctx context.Context,
	id int) (model.Project, error) {
	var project model.Project
	err := r.db.QueryRow(ctx,
		`SELECT id,user_id,name
	FROM projects
	WHERE id = $1`,
		id).Scan(&project.ID, &project.UserID, &project.Name)
	return project, err
}

func (r *ProjectRepository) GetByUserID(ctx context.Context,
	user_id int) ([]model.Project, error) {
	rows, err := r.db.Query(ctx,
		`SELECT id,user_id,name
	FROM projects
	WHERE user_id = $1`, user_id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var projects []model.Project

	for rows.Next() {
		var project model.Project

		err := rows.Scan(
			&project.ID,
			&project.UserID,
			&project.Name,
		)
		if err != nil {
			return nil, err
		}
		projects = append(projects, project)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return projects, err
}

func (r *ProjectRepository) Update(ctx context.Context,
	project model.Project) (model.Project, error) {
	err := r.db.QueryRow(ctx,
		`UPDATE projects
	SET name = $1
	WHERE id = $2
	RETURNING id,user_id,name`, project.Name, project.ID).Scan(
		&project.ID,
		&project.UserID,
		&project.Name,
	)
	if err != nil {
		return model.Project{}, err
	}
	return project, nil
}

func (r *ProjectRepository) Delete(ctx context.Context,
	id int) error {
	result, err := r.db.Exec(ctx,
		`DELETE FROM projects
		WHERE id = $1`,
		id)
	if err != nil {
		return err
	}
	if result.RowsAffected() == 0 {
		return errors.New("project not found")
	}
	return nil
}
