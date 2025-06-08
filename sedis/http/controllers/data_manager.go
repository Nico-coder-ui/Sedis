package controllers

import (
	"net/http"
	"sedis/stokage"

	"github.com/gin-gonic/gin"
)

type Answer struct {
	DATA string `json:"id"`
}

func TtlHandler(store *stokage.Store) gin.HandlerFunc {
	return func(c *gin.Context) {
		Ttl(c, store)
	}
}

func Ttl(c *gin.Context, store *stokage.Store) {
	key := c.GetString("key")
	if key != "" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "key error"})
		return
	}

	var a Answer
	a.DATA = store.Ttl(key)
	c.JSON(http.StatusOK, a)
}

func List(c *gin.Context) {

}

func Set(c *gin.Context) {

}

func GetHandler(store *stokage.Store) gin.HandlerFunc {
	return func(c *gin.Context) {
		Get(c, store)
	}
}

func Get(c *gin.Context, store *stokage.Store) {
	key := c.GetString("key")
	if key != "" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "key error"})
		return
	}

	var a Answer
	a.DATA, _ = store.Get(key)
	c.JSON(http.StatusOK, a)
}

func Del(c *gin.Context) {

}

func Flushall(c *gin.Context) {

}
