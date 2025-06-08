package http

import (
	"net/http"
	"sedis/http/controllers"
	"sedis/stokage"

	"github.com/gin-gonic/gin"
)

func SetupRouter(store *stokage.Store) *gin.Engine {
	r := gin.Default()

	// Middleware CORS
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	})

	r.POST("/ttl", controllers.TtlHandler(store))
	r.POST("/set", controllers.Set)
	r.POST("/del", controllers.Del)
	r.POST("/flushall", controllers.Flushall)
	r.POST("/save", controllers.Save)
	r.POST("/load", controllers.Load)

	r.GET("/exists", controllers.Exists)
	r.GET("/list", controllers.List)
	r.GET("/get", controllers.GetHandler(store))

	return r
}
