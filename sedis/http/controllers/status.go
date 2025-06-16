package controllers

import (
	"net/http"
	"sedis/stokage"

	"github.com/gin-gonic/gin"
)

func ExistsHandler(store *stokage.Store) gin.HandlerFunc {
	return func(c *gin.Context) {
		Exists(c, store)
	}
}

func Exists(c *gin.Context, store *stokage.Store) {
	key := c.GetString("key")
	if key != "" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "key error"})
		return
	}

	var a Answer
	a.DATA = store.Exists(key)
	c.JSON(http.StatusOK, a)
}
