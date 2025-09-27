package dto

import "schedulr/internal/models"

func ToUserResponse(u *models.User) UserResponse {
	return UserResponse{
		ID:    u.ID,
		Name:  u.Name,
		Email: u.Email,
	}
}
