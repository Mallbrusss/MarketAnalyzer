package handlers

import "github.com/labstack/echo/v4"

type InstrumentHandler struct {
}

func NewInstrumentHandler() *InstrumentHandler {
	return &InstrumentHandler{}
}

func (i *InstrumentHandler) GetInstrumentUID(c echo.Context) error {

	return nil
}

func (i *InstrumentHandler) GetCandles(c echo.Context) error {
	return nil
}
