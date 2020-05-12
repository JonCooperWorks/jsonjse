package jsonjse

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type JSEHandlers struct {
	*ServerConfig
}

func (h *JSEHandlers) HandleTodaysPricesLookup(c *gin.Context) {
	symbols, err := h.JSE.GetPricesForDate(time.Now())
	if err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, &symbols)
}
