package dtos 

type LoginReq struct{
	Email string  `json:"email"`
	Contact_number string `json:"contact_number"`
	Password string `json:"password"`
} 

type LoginRes struct{
	Token string
}