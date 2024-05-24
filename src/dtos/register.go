package dtos

type RegisterReq struct {
	Username       string `json:"username"`
	Email          string `json:"email"`
	Gender         string `json:"gender"`
	Contact_Number string `json:"contact_number"`
	Password       string `json:"password"`
	Role           string `json:"role_id"`
}
