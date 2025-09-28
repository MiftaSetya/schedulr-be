package dto

type CreateListInput struct {
	Name     string `json:"name" binding:"required"`
	BoardID  uint   `json:"board_id" binding:"required"`
	Position int    `json:"position,omitempty"`
}

type ListResponse struct {
	ID       uint           `json:"id"`
	Name     string         `json:"name"`
	BoardID  uint           `json:"board_id"`
	Position int            `json:"position"`
	Tasks    []TaskResponse `json:"tasks,omitempty"`
}
