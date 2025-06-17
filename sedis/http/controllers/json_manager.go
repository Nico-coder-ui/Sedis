package controllers

import (
	"io"
	"net/http"
	"sedis/stokage"
	"strings"

	"github.com/gin-gonic/gin"
)

func SaveHandler(store *stokage.Store) gin.HandlerFunc {
	return func(c *gin.Context) {
		Save(c, store)
	}
}

func Save(c *gin.Context, store *stokage.Store) {
	fileName := "json/store.json"
	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.String(http.StatusBadRequest, "failed to read body")
		return
	}
	msg := string(body)
	tokens := strings.Fields(msg)

	if len(tokens) >= 2 {
		fileName = "json/" + tokens[1]
		if !strings.HasSuffix(fileName, ".json") {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "File name must be a .json\n"})
			return
		}
	}

	ok := store.Save(fileName)
	if ok {
		c.JSON(http.StatusOK, nil)
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save data\n"})
	}
}

func LoadHandler(store *stokage.Store) gin.HandlerFunc {
	return func(c *gin.Context) {
		Load(c, store)
	}
}

func Load(c *gin.Context, store *stokage.Store) {
	fileName := "json/store.json"
	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.String(http.StatusBadRequest, "failed to read body")
		return
	}
	msg := string(body)
	tokens := strings.Fields(msg)

	if len(tokens) >= 2 {
		fileName = "json/" + tokens[1]
		if !strings.HasSuffix(fileName, ".json") {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "File name must be a .json\n"})
			return
		}
	}

	ok := store.Load(fileName)
	if ok {
		c.JSON(http.StatusOK, nil)
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save data\n"})
	}
}
