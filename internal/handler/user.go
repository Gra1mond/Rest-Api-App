package handler

import (
	"encoding/json"
	"net/http"
	"rest_api/internal/service"
)

type UserHandler struct{
	service *service.UserService
}

type RegisterRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func NewUserHandler (service *service.UserService)*UserHandler{
	return &UserHandler{
		service: service,
	}
}

func (h *UserHandler)Create(w http.ResponseWriter, r *http.Request){
	var req RegisterRequest

	err:=json.NewDecoder(r.Body).Decode(&req)

	if err!=nil{
		http.Error(w,"invalid request body",http.StatusBadRequest)
		return
	}

	user,err:=h.service.Register(
		r.Context(),
		req.Email,
		req.Password,
	)
	if err!=nil{
		http.Error(w,err.Error(),http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}