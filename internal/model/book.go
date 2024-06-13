package model

type BookingRequest struct {
	UserID   int   `json:"user_id" binding:"required"`
	RouteIDs []int `json:"route_ids" binding:"required"`
}

type BookingResponse struct {
	Success   bool   `json:"success"`
	FailedIDs []int  `json:"failed_ids"`
	Message   string `json:"message"`
}
