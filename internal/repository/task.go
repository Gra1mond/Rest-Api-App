// package repository

// import (
// 	"errors"
// 	"rest_api/internal/model"
// )

// type TaskRepository struct{
// 	tasks []model.Task
// 	nextID int
// }

// func NewTaskRepository()*TaskRepository{
// 	return  &TaskRepository{
// 		tasks: make([]model.Task, 0),
// 		nextID: 1,
// 	}
// }

// func (r *TaskRepository) Create(task model.Task)(model.Task,error){
// 	task.ID=r.nextID

// 	r.nextID++

// 	r.tasks=append(r.tasks, task)

// 	return task,nil
// }

// func (r *TaskRepository) GetByID(id int)(model.Task,error) {
// 	for _,task:=range r.tasks{
// 		if task.ID == id {
// 			return task, nil
// 		}
// 	}
// 	return model.Task{},errors.New("failed to find id")
// }

// func (r *TaskRepository) GetAll() ([]model.Task, error) {
//     return r.tasks, nil
// }

//postgresql

package repository

import (
	"context"
	"rest_api/internal/model"

	"github.com/jackc/pgx/v5"
)

type TaskRepository struct{
	db *pgx.Conn
}

func NewTaskRepository(db *pgx.Conn) *TaskRepository{
	return &TaskRepository{
		db:db,
	}
}

func  (r *TaskRepository) GetBy(
	ctx context.Context,
	id int)(model.Task,error){
		var task model.Task
		err:=r.db.QueryRow(ctx,`SELECT id,project_id,description,done
		FROM tasks
		where id = $1`,id).Scan(&task.ID,
        &task.ProjectID,
        &task.Title,
        &task.Description,
        &task.Done,)

		return task,err
}

func (r* TaskRepository)GetAll(ctx context.Context)([]model.Task,error){
	rows,err:=r.db.Query(ctx,
	`SELECT id,project_id,title,description,done
	FROM tasks`)
	if err!=nil{
		return nil,err
	}
	defer rows.Close()

	var tasks []model.Task

	for rows.Next(){
		var task model.Task
		if err:=rows.Scan(
		&task.ID,
        &task.ProjectID,
        &task.Title,
        &task.Description,
        &task.Done,
		);err!=nil{
			return nil,err
		}
		tasks = append(tasks,task)
	}
	if err:=rows.Err();err!=nil{
		return nil,err
	}
	return tasks,nil
}


func (r *TaskRepository)Create(
	ctx context.Context,task model.Task,
	)(model.Task,error){
		err:=r.db.QueryRow(ctx,
		`INSERT INTO TASKS(project_id,title,description,done)
		VALUES($1,$2,$3,$4)
		RETURNING id`,
	task.ProjectID,
task.Title,
task.Description,
task.Done).Scan(&task.ID)
return task,err
}