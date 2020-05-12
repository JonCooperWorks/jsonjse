package jsonjse

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type JSEHandlers struct {
	*ServerConfig
}

func (h *JSEHandlers) HandleTodaysPricesLookup(c *gin.Context) {
	symbols, err := h.JSE.GetTodaysPrices()
	if err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, &symbols)
}

func (h *JSEHandlers) HandleNewsArticle(c *gin.Context) {
	newsArticles, err := GetTodaysNews()
	if err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, &newsArticles)
}
