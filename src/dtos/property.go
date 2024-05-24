package dtos

type PropertyReq struct{
	Id string `json:"id"`
	Name string 	 `json:"name"`
	Price int	 `json:"price"`
	Location string  `json:"location"`
	User string  `json:"user_id"`
	Category string  `json:"category_id"`
	IsApproved bool  `json:"is_approved"`
	IsBooked bool `json:"is_booked"`
}
