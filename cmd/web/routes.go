package main

import (
	"github.com/labstack/echo/v4"
	"github.com/mariogarzac/snowflake_demo/internal/handlers"
)

func routes(e *echo.Echo, h *handlers.Handlers) {
	e.GET("/", h.GetAllRecords)
}
