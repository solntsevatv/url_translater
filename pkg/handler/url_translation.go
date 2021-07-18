package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/solntsevatv/url_translater/internal/url_translater"
)

func (h *Handler) longToShort(c *gin.Context) {
	var input url_translater.LongURL

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	short_url, err := h.services.UrlTranslation.CreateShortURL(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"url": short_url,
	})
}

func (h *Handler) ShortToLong(c *gin.Context) {
	input := url_translater.ShortURL{Id: 1, LinkUrl: ""}

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	long_url, err := h.services.UrlTranslation.GetLongURL(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"url": long_url,
	})
}
