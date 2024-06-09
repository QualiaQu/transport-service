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

func MapPGToHandler(pgTransports []TransportPG) []TransportHandler {
	handlers := make([]TransportHandler, len(pgTransports))
	for i, pg := range pgTransports {
		handlers[i] = pg.ToHandler()
	}

	return handlers
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

func MapHandlerToPG(handlerTransports []TransportHandler) []TransportPG {
	pgTransports := make([]TransportPG, len(handlerTransports))
	for i, handler := range handlerTransports {
		pgTransports[i] = handler.ToPG()
	}

	return pgTransports
}
