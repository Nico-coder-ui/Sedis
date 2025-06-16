package controllers

import (
	"net/http"
	"sedis/stokage"

	"github.com/gin-gonic/gin"
)

func SaveHandler(store *stokage.Store) gin.HandlerFunc {
	return func(c *gin.Context) {
		Save(c, store)
	}
}

func Save(c *gin.Context, store *stokage.Store) {
	filename := c.GetString("filename")
	if filename != "" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "filename error"})
		return
	}

	isAccepted := store.Save(filename)
	if isAccepted {
		c.JSON(http.StatusOK, nil)
	} else {
		c.JSON(http.StatusNotAcceptable, nil)
	}
}

func LoadHandler(store *stokage.Store) gin.HandlerFunc {
	return func(c *gin.Context) {
		Load(c, store)
	}
}

func Load(c *gin.Context, store *stokage.Store) {
	filename := c.GetString("filename")
	if filename != "" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "filename error"})
		return
	}

	isAccepted := store.Load(filename)
	if isAccepted {
		c.JSON(http.StatusOK, nil)
	} else {
		c.JSON(http.StatusNotAcceptable, nil)
	}
}
