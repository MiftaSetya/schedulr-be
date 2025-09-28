package dto

type CreateBoardInput struct {
	Name    string `json:"name" binding:"required"`
	OwnerID uint   `json:"owner_id" binding:"required"`
}

type BoardResponse struct {
	ID      uint   `json:"id"`
	Name    string `json:"name"`
	OwnerID uint   `json:"owner_id"`
}
