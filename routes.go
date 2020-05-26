package jsonjse

import "github.com/gin-gonic/gin"

func Router(config *ServerConfig) *gin.Engine {
	router := gin.Default()
	router.Use(func (c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
	})

	handlers := &JSEHandlers{ServerConfig: config}
	stocks := router.Group("/jse")
	stocks.GET("/today", handlers.HandleTodaysPricesLookup)
	stocks.GET("/news", handlers.HandleNewsArticle)
	return router
}
