package handler

import (
	"encoding/json"
	"net/http"
	"rest_api/internal/midleware"
	"rest_api/internal/model"
	"rest_api/internal/service"
	"strconv"
)

type TaskHandler struct{
	service *service.TaskService
}

func NewTaskHandler(service *service.TaskService)*TaskHandler{
	return &TaskHandler{
		service: service,
	}
}

func (h *TaskHandler)Create(w http.ResponseWriter, r *http.Request){
	projectIDStr:=r.PathValue("projectID")

	projectID,err:=strconv.Atoi(projectIDStr)

	if err!=nil{
		http.Error(w,"invalid project id",http.StatusBadRequest)
		return
	}

	var task model.Task

	if err:=json.NewDecoder(r.Body).Decode(&task);err!=nil{
		http.Error(w,"invalid body",http.StatusBadRequest)
		return
	}
	userID,ok:=midleware.GetUserID(r.Context())
	if !ok{
		http.Error(w,"unuathorized",http.StatusUnauthorized)
		return
	}
	createdTask,err:=h.service.Create(r.Context(),
	userID,
	task,
	projectID)
	if err!=nil{
		http.Error(w,err.Error(),http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdTask)
}

func (h *TaskHandler)GetByProjectID(w http.ResponseWriter,r *http.Request){
	projectIDStr:=r.PathValue("projectID")

	projectID,err:=strconv.Atoi(projectIDStr)
	if err!=nil{
		http.Error(w,"invalid project id",http.StatusBadRequest)
		return
	}
	userID,ok:=midleware.GetUserID(r.Context())
	if !ok{
		http.Error(w,"unuathorized",http.StatusUnauthorized)
		return
	}

	task,err:=h.service.GetByProjectID(r.Context(),userID,projectID)
	if err!=nil{
		http.Error(w,err.Error(),http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(task)
}

func (h *TaskHandler)GetByID(w http.ResponseWriter, r *http.Request){
	taskIDStr:=r.PathValue("id")
	taskID,err:=strconv.Atoi(taskIDStr)

	if err!=nil{
		http.Error(w,"invalid id",http.StatusBadRequest)
		return
	}
	userID,ok:=midleware.GetUserID(r.Context())
	if !ok{
		http.Error(w,"unuathorized",http.StatusUnauthorized)
		return
	}

	task,err:=h.service.GetByID(r.Context(),userID,taskID)

	if err!=nil{
		http.Error(w,err.Error(),http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(task)
}

func(h *TaskHandler)Update(w http.ResponseWriter, r *http.Request){
	var req model.Task

	err:=json.NewDecoder(r.Body).Decode(&req)

	if err!=nil{
		http.Error(w,"invalid request body",http.StatusBadRequest)
		return
	}
	userID,ok:=midleware.GetUserID(r.Context())
	if !ok{
		http.Error(w,"unuathorized",http.StatusUnauthorized)
		return
	}

	task,err:=h.service.Update(r.Context(),userID,req)
	if err!=nil{
		http.Error(w,err.Error(),http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(task)
}

func (h *TaskHandler)Delete(w http.ResponseWriter, r *http.Request){
	taskIDStr:=r.PathValue("id")

	taskID,err:=strconv.Atoi(taskIDStr)
	if err!=nil{
		http.Error(w,"invalid id",http.StatusBadRequest)
		return
	}
	userID,ok:=midleware.GetUserID(r.Context())
	if !ok{
		http.Error(w,"unuathorized",http.StatusUnauthorized)
		return
	}

	deleteError:=h.service.Delete(r.Context(),userID,taskID)
	if deleteError!=nil{
		http.Error(w,deleteError.Error(),http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusAccepted)
}