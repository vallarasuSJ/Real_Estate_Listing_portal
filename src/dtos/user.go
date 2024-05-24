package dtos

import "real_estate/src/database/models"

type Users struct {
	Id             string `json:"id"`
	Email          string `json:"email"`
	Contact_number string `json:"contact_number"`
	Username       string `json:"username"`
	RoleId          string `json:"role_id"`
	RoleName       string `json:"role_name"`
	Gender         string `json:"gender"`
}

type UserReq struct {
	*models.Users
	RoleName string `json:"role_name"`
}
