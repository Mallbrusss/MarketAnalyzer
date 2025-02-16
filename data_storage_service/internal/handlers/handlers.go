package handlers

import (
	"data-storage/internal/models"
	"fmt"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
)

// Repository Запросы на создание происходит через кафку
type Repository interface {
	GetInstrumentUIDAndFigi(ticker string) (models.Ids, error)
	GetCandles(instrumentUID string) ([]models.HistoricCandle, error)
}

type Handler struct {
	Repository Repository
}

func NewHandler(repository Repository) *Handler {
	return &Handler{
		Repository: repository,
	}
}

func (h *Handler) GetInstrumentUIDAndFigi(c echo.Context) error {
	c.Request().Header.Set("Content-Type", "application/json")
	ticker := c.QueryParam("ticker")

	instrumentIDS, err := h.Repository.GetInstrumentUIDAndFigi(ticker)
	if err != nil {
		log.Printf("Error getting instrument uid: %v", err)
		return c.JSON(http.StatusNotFound, echo.Map{
			"error": "Invalid name, not found instrument uid",
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"instrumentIDS": instrumentIDS,
	})
}

func (h *Handler) GetCandles(c echo.Context) error {
	c.Request().Header.Set("Content-Type", "application/json")

	//var historicCandle []models.HistoricCandle
	historicCandle, err := h.Repository.GetCandles(c.QueryParam("instrument_id"))
	fmt.Println("err", err)
	if err != nil {
		fmt.Println("err:", err)
		return c.JSON(http.StatusNotFound, echo.Map{
			"error": "Invalid uid, not found candles",
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"candles": historicCandle,
	})
}
