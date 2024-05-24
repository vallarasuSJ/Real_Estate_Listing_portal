package dtos

type BookingReq struct{
	Id string `json:"id"`
	User string  `json:"user_id"`
	Property string  `json:"property_id"`
}

