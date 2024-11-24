package handlers

import (
	"net/http"
	"tinkoff-api/internal/models"

	"github.com/labstack/echo/v4"
)

type StockExchange interface {
	GetClosePrices(instruments []string) ([]models.ClosePrice, error)
	GetAllBonds(instrumentStatus string) ([]models.PlacementPrice, error)
	GetCandles(instrumentInfo map[string]any) ([]models.HistoricCandle, error)
}

type Handler struct {
	Service StockExchange
}

func NewHandler(service StockExchange) *Handler {
	return &Handler{
		Service: service,
	}
}

func (h *Handler) GetClosePricesHandler(c echo.Context) error {
	c.Request().Header.Set("Content-Type", "application/json")

	var req models.GetClosePricesRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": "Invalid request format",
		})
	}

	instruments := make([]string, 0, len(req.Instruments))
	for _, instrument := range req.Instruments {
		instruments = append(instruments, instrument.InstrumentID)
	}

	closePrices, err := h.Service.GetClosePrices(instruments)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": "Failed to fetch close prices",
		})
	}

	response := make(map[string]models.ClosePrice, len(closePrices))
	for _, v := range closePrices {
		response[v.InstrumentUid] = v
	}

	return c.JSON(http.StatusOK, response)
}

func (h *Handler) GetAllBonds(c echo.Context) error {
	c.Request().Header.Set("Content-Type", "application/json")

	//TODO: Change this func param
	allBonds, err := h.Service.GetAllBonds("INSTRUMENT_STATUS_BASE")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": "Failed to fetch all bonds",
		})
	}

	return c.JSON(http.StatusOK, allBonds)
}

func (h *Handler) GetCandles(c echo.Context) error {
	c.Request().Header.Set("Content-Type", "application/json")
	var req models.GetCandlesRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": "Invalid request format",
		})
	}

	instumentInfo := map[string]any{
		"figi":         req.Figi,
		"from":         req.From,
		"to":           req.To,
		"interval":     req.Interval,
		"instrumentId": req.InstrumentId,
	}

	candles, err := h.Service.GetCandles(instumentInfo)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": "Failed to fetch all candles",
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		req.InstrumentId: candles,
	})
}
