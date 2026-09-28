package main

import (
	"net/http"
	"rest_api/internal/handler"
	"rest_api/internal/midleware"
	"rest_api/internal/repository"
	"rest_api/internal/service"
)

func main(){
	mux:=http.NewServeMux()

	repo:=repository.NewTaskRepository()
	service:=service.NewService(repo)
	handler:=handler.NewHandler(service)

	mux.HandleFunc("GET /tasks", handler.GetAll)
	mux.HandleFunc("GET /tasks/{id}", handler.GetByID)
	mux.Handle(
        "POST /tasks",
        midleware.Auth(
            http.HandlerFunc(handler.Create),
        ),
    )

    
    app := midleware.Logger(mux)

	http.ListenAndServe(":8080",app)

}