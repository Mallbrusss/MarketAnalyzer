package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Handlers struct {
	Ingestion Ingestion
}

func NewHandlers(ingestion Ingestion) *Handlers{
	return &Handlers{
		Ingestion: ingestion,
	}
}

type Ingestion interface {
	Ingest(dataURL string) error
}

func (h *Handlers) Ingest(c echo.Context) error {
	req := struct {
		DataURL string `json:"data_url"`
	}{}

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})
	}

	if err := h.Ingestion.Ingest(req.DataURL); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]string{"status": "File ingestion started"})
}
