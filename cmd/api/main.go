package main

import (
	"context"
	"log"
	"net/http"
	"rest_api/internal/database"
	"rest_api/internal/handler"
	"rest_api/internal/midleware"
	"rest_api/internal/repository"
	"rest_api/internal/service"
)

func main(){
	mux:=http.NewServeMux()
	ctx:=context.Background()

	conn,err:=database.Connect(ctx,
		"postgres://postgres:1@localhost:5432/taskmanager",
	)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close(ctx)


	//
	userRepo := repository.NewUserRepository(conn)
	projectRepo := repository.NewProjectRepository(conn)
	taskRepo := repository.NewTaskRepository(conn)
	//

	//
	userService:=service.NewUserService(userRepo)
	projectService:=service.NewProjectService(projectRepo)
	taskService:=service.NewTaskService(
		taskRepo,
		projectRepo,
	)
	//

	//
	userHandler:=handler.NewUserHandler(userService)
	projectHandler:=handler.NewProjectHandler(projectService)
	taskHandler:=handler.NewTaskHandler(taskService)
	//

	//
	mux.HandleFunc(
		"POST /register",
		userHandler.Create,
	)

	mux.HandleFunc(
		"POST /login",
		userHandler.Login,
	)
	//

	//
	mux.Handle(
		"POST /projects",
		midleware.Auth(http.HandlerFunc(projectHandler.Create)),
	)

	mux.Handle(
		"GET /projects",
		midleware.Auth(http.HandlerFunc(projectHandler.GetByUserID)),
	)

	mux.Handle(
		"GET /projects/{id}",
		midleware.Auth(http.HandlerFunc(projectHandler.GetByID)),
	)

	mux.Handle(
		"PUT /projects/{id}",
		midleware.Auth(http.HandlerFunc(projectHandler.Update)),
	)

	mux.Handle(
		"DELETE /projects/{id}",
		midleware.Auth(http.HandlerFunc(projectHandler.Delete)),
	)
	//


	//
	mux.Handle("POST /projects/{projectID}/tasks",
	midleware.Auth(
		http.HandlerFunc(taskHandler.Create),
))
mux.Handle(
		"GET /projects/{projectID}/tasks",
		midleware.Auth(
			http.HandlerFunc(taskHandler.GetByProjectID),
		),
	)

	mux.Handle(
		"GET /tasks/{id}",
		midleware.Auth(
			http.HandlerFunc(taskHandler.GetByID),
		),
	)

	mux.Handle(
		"PUT /tasks/{id}",
		midleware.Auth(
			http.HandlerFunc(taskHandler.Update),
		),
	)

	mux.Handle(
		"DELETE /tasks/{id}",
		midleware.Auth(
			http.HandlerFunc(taskHandler.Delete),
		),
	)
    //
    app := midleware.Logger(mux)

	err=http.ListenAndServe(":8080",app)
	if err != nil {
		log.Fatal(err)
	}

}