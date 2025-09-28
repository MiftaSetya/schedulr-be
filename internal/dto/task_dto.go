package dto

import "time"

type CreateBoardTaskInput struct {
	Title    string `json:"title" binding:"required"`
	ListID   uint   `json:"list_id" binding:"required"`
	Position int    `json:"position,omitempty"`
}

type CreateReminderTaskInput struct {
	Title       string     `json:"title" binding:"required"`
	Description string     `json:"description,omitempty"`
	DueDate     *time.Time `json:"due_date" binding:"required"`
	UserID      uint       `json:"user_id" binding:"required"`
}

type TaskResponse struct {
	ID          uint       `json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description,omitempty"`
	Completed   bool       `json:"completed"`
	DueDate     *time.Time `json:"due_date,omitempty"`
	ListID      *uint      `json:"list_id,omitempty"`
	Position    int        `json:"position,omitempty"`
	UserID      *uint      `json:"user_id,omitempty"`
}
