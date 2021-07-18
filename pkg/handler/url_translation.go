package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
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

	logrus.Info("short_url=", short_url, " was added in db")
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

	logrus.Info("long_url=", long_url, " was gotten from db")
}
