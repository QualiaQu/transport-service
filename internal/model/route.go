package model

type RouteRequest struct {
	Origin             string `json:"origin" binding:"required"`
	Destination        string `json:"destination" binding:"required"`
	Date               string `json:"date" binding:"required"`
	PreferredTransport []int  `json:"preferred_transport" binding:"required"`
}

type RouteResponse struct {
	ID                int     `json:"id"`
	TransportType     int     `json:"transport_type"`
	Price             float64 `json:"price"`
	DepartureDatetime string  `json:"departure_datetime"`
	ArrivalDatetime   string  `json:"arrival_datetime"`
}

type RoutePG struct {
	ID                int     `db:"id"`
	TransportType     int     `db:"transport_type"`
	Price             float64 `db:"price"`
	DepartureDatetime string  `db:"departure_datetime"`
	ArrivalDatetime   string  `db:"arrival_datetime"`
}

func (pg *RoutePG) ToResponse() RouteResponse {
	return RouteResponse{
		ID:                pg.ID,
		TransportType:     pg.TransportType,
		Price:             pg.Price,
		DepartureDatetime: pg.DepartureDatetime,
		ArrivalDatetime:   pg.ArrivalDatetime,
	}
}
