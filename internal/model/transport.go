package model

type TransportPG struct {
	ID   int    `pg:"id"`
	Name string `pg:"name"`
}

func (pg *TransportPG) ToHandler() TransportHandler {
	return TransportHandler{
		ID:   pg.ID,
		Name: pg.Name,
	}
}

type TransportHandler struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func (handler *TransportHandler) ToPG() TransportPG {
	return TransportPG{
		ID:   handler.ID,
		Name: handler.Name,
	}
}
