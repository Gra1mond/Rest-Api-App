package model


type Project struct {
	ID     int    `json:"id"`
	UserID int    `json:"user_id"`
	Name   string `json:"name"`
}