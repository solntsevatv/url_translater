package handler

import (
	"github.com/solntsevatv/url_translater/pkg/service"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	short_url := router.Group("/short")
	{
		short_url.POST("", h.longToShort)
	}

	long := router.Group("/long")
	{
		long.POST("", h.ShortToLong)
	}

	return router
}
