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
	today := time.Now().Format("2006-01-02")
	symbols, err := h.JSE.DailyPrices(today)
	if err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, &symbols)
}

func (h *JSEHandlers) HandleNewsArticle(c *gin.Context) {
	today := time.Now().Format("2006-01-02")
	newsArticles, err := h.JSE.DailyNews(today)
	if err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, &newsArticles)
}
