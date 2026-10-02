package handler

import (
	"encoding/json"
	"net/http"
	"rest_api/internal/model"
	"rest_api/internal/service"
	"strconv"
)

type ProjectHandler struct{
	service *service.ProjectService
}

func NewProjectHandler(service *service.ProjectService)*ProjectHandler{
	return &ProjectHandler{
		service: service,
	}
}

func (h *ProjectHandler)Create(w http.ResponseWriter, r *http.Request){
	var req model.Project

	err:=json.NewDecoder(r.Body).Decode(&req)

	if err!=nil{
		http.Error(w,"invalid request body",http.StatusBadRequest)
		return
	}

	userID := 1// достаём из context после Auth middleware

	project, err := h.service.Create(
		r.Context(),
		userID,
		req,
	)

	if err!=nil{
		http.Error(w,err.Error(),http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(project)
}

func(h *ProjectHandler)GetByID(w http.ResponseWriter, r *http.Request){
	ProjectIDStr:=r.PathValue("id")

	projectID,err:=strconv.Atoi(ProjectIDStr)

	if err!=nil{
		http.Error(w,"project id is required",http.StatusBadRequest)
		return
	}

	userID := 1// из r.Context()

	project, err := h.service.GetByID(
		r.Context(),
		userID,
		projectID,
	)
	if err!=nil{
		http.Error(w,err.Error(),http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(project)
}

func (h *ProjectHandler) GetByUserID(
	w http.ResponseWriter,
	r *http.Request,
) {
	userID :=1 // достаём из r.Context()

	projects, err := h.service.GetByUserID(
		r.Context(),
		userID,
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(projects)
}

func(h *ProjectHandler)Update(w http.ResponseWriter, r *http.Request){
	var req model.Project

	projectIDStr:=r.PathValue("id")

	projectID, err := strconv.Atoi(projectIDStr)

	if err!=nil{
		http.Error(w,"invalid project id",http.StatusBadRequest)
		return
	}

	if err:=json.NewDecoder(r.Body).Decode(&req);err!=nil{
		http.Error(w,"invalid request body",http.StatusBadRequest)
		return
	}

	req.ID=projectID

	userID:=1//из r.Contex()

	project,err:=h.service.Update(r.Context(),
		userID,
		req,)
	if err!=nil{
		http.Error(w,err.Error(),http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(project)
}

func (h *ProjectHandler) Delete(
	w http.ResponseWriter,
	r *http.Request,
) {
	projectIDStr := r.PathValue("id")

	projectID, err := strconv.Atoi(projectIDStr)
	if err != nil {
		http.Error(w, "invalid project id", http.StatusBadRequest)
		return
	}

	userID :=1 // из r.Context()

	err = h.service.Delete(
		r.Context(),
		userID,
		projectID,
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}