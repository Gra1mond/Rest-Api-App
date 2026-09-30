package handler

import (
	"encoding/json"
	"net/http"
	"rest_api/internal/model"
	"rest_api/internal/service"
	"strconv"
)

type TaskHandler struct {
	service *service.TaskService
}

func NewHandler(service *service.TaskService)*TaskHandler{
	return &TaskHandler{
		service: service,
	}
}

func (h *TaskHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	task, err := h.service.GetByID(id)

	if err != nil {
		http.Error(w, "task not found", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(task)
}

func (h *TaskHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	tasks, err := h.service.GetAll()
	if err != nil {
		http.Error(w, "failed to get tasks", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}

func (h *TaskHandler) Create(w http.ResponseWriter, r *http.Request) {
	var task model.Task
	err := json.NewDecoder(r.Body).Decode(&task)

	if err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}
	createdTask, err := h.service.Create(task)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdTask)
}
