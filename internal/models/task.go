package models

import (
	"time"

	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	Title       string     `gorm:"size:255;not null" json:"title"`
	Description string     `gorm:"type:text" json:"description,omitempty"`
	Completed   bool       `gorm:"default:false" json:"completed"`
	DueDate     *time.Time `json:"due_date,omitempty"`
	UserID      *uint      `json:"user_id,omitempty"`
	ListID      *uint      `json:"list_id,omitempty"`
	Position    int        `json:"position,omitempty"`
}

type List struct {
	gorm.Model
	Name     string `gorm:"size:255;not null" json:"name"`
	BoardID  uint   `json:"board_id"`
	Position int    `json:"position"`
	Tasks    []Task `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"tasks,omitempty"`
}

type Board struct {
	gorm.Model
	Name    string `gorm:"size:255;not null" json:"name"`
	OwnerID uint   `json:"owner_id"`
	Lists   []List `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"lists,omitempty"`
}
