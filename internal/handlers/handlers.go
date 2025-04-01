package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/mariogarzac/snowflake_demo/pkg/config"
)

var handlers *Handlers

type Handlers struct {
	AppConfig *config.Config
}

func NewRepo(a *config.Config) *Handlers {
	return &Handlers{
		AppConfig: a,
	}
}

func (h *Handlers) NewRepository(ha *Handlers) {
	handlers = ha
}

func (h *Handlers) GetAllRecords(c echo.Context) {
	handlers.AppConfig.Database.GetAllRecords()
}
