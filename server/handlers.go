package server

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (w *TinyWatcherServer) handleWatchAddress(c echo.Context) error {
	address := c.Param("address")
	if address == "" {
		return c.JSON(http.StatusBadRequest, "address is required")
	}

	err := w.processWatchAddress(address)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusCreated, "success")
}

func (w *TinyWatcherServer) handleWatchState(c echo.Context) error {
	response := w.processWatchState()
	return c.JSON(http.StatusOK, response)
}
