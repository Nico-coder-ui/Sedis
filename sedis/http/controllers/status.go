package controllers

import (
	"io"
	"net/http"
	"sedis/stokage"
	"strings"

	"github.com/gin-gonic/gin"
)

func ExistsHandler(store *stokage.Store) gin.HandlerFunc {
	return func(c *gin.Context) {
		Exists(c, store)
	}
}

func Exists(c *gin.Context, store *stokage.Store) {
	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.String(http.StatusBadRequest, "failed to read body")
		return
	}
	msg := string(body)
	tokens := strings.Fields(msg)

	key := tokens[1]
	value := store.Exists(key)

	var a Answer
	a.DATA = value
	c.JSON(http.StatusOK, a)
}
