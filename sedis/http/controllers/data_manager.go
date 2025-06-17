package controllers

import (
	"io"
	"net/http"
	"sedis/stokage"
	"strconv"
	"strings"
	"time"

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
	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.String(http.StatusBadRequest, "failed to read body")
		return
	}
	msg := string(body)
	tokens := strings.Fields(msg)

	if len(tokens) < 2 {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Usage: TTL key\n"})
		return
	}

	key := tokens[1]

	value := store.Ttl(key)

	var a Answer
	a.DATA = value
	c.JSON(http.StatusOK, a)
}

func ListHandler(store *stokage.Store) gin.HandlerFunc {
	return func(c *gin.Context) {
		List(c, store)
	}
}

func List(c *gin.Context, store *stokage.Store) {
	result := store.List()
	if result != "" {
		var a Answer
		a.DATA = result
		c.JSON(http.StatusOK, a)
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No data\n"})
	}
}

func SetHandler(store *stokage.Store) gin.HandlerFunc {
	return func(c *gin.Context) {
		Set(c, store)
	}
}

func Set(c *gin.Context, store *stokage.Store) {
	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.String(http.StatusBadRequest, "failed to read body")
		return
	}
	msg := string(body)
	tokens := strings.Fields(msg)

	if len(tokens) == 0 {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Empty command\n"})
		return
	}

	if len(tokens) < 3 {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Usage: SET key value [NX|XX] [EX seconds]\n"})
		return
	}

	key := tokens[1]
	value := tokens[2]

	var expireInSec string
	var hasExpire bool

	for i := 3; i < len(tokens); i++ {
		switch strings.ToUpper(tokens[i]) {
		case "EX":
			if i+1 < len(tokens) {
				expireInSec = tokens[i+1]
				hasExpire = true
				i++
			} else {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Missing value for EX\n"})
				return
			}

		case "NX":
			if store.Exists(key) != "0" {
				c.JSON(http.StatusOK, nil)
				return
			}

		case "XX":
			if store.Exists(key) == "0" {
				c.JSON(http.StatusOK, nil)
				return
			}

		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Unknown modifier: " + tokens[i] + "\n"})
			return
		}
	}

	if hasExpire {
		seconds, err := strconv.Atoi(expireInSec)
		if err != nil {
			c.JSON(http.StatusBadRequest, nil)
			return
		}
		expireAt := time.Now().Add(time.Duration(seconds) * time.Second)
		store.Expire(key, expireAt)
	}

	store.Set(key, value)
	c.JSON(http.StatusOK, nil)
}

func GetHandler(store *stokage.Store) gin.HandlerFunc {
	return func(c *gin.Context) {
		Get(c, store)
	}
}

func Get(c *gin.Context, store *stokage.Store) {
	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.String(http.StatusBadRequest, "failed to read body")
		return
	}
	msg := string(body)
	tokens := strings.Fields(msg)

	if len(tokens) < 2 {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Usage: GET key\n"})
		return
	}
	key := tokens[1]

	value, ok := store.Get(key)
	if ok {
		var a Answer
		a.DATA = value
		c.JSON(http.StatusOK, a)
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Key not found\n"})
	}
}

func DelHandler(store *stokage.Store) gin.HandlerFunc {
	return func(c *gin.Context) {
		Del(c, store)
	}
}

func Del(c *gin.Context, store *stokage.Store) {
	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.String(http.StatusBadRequest, "failed to read body")
		return
	}
	msg := string(body)
	tokens := strings.Fields(msg)

	if len(tokens) < 2 {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Usage: DEL key\n"})
		return
	}
	key := tokens[1]
	b := store.Del(key)

	if b {
		c.JSON(http.StatusOK, nil)
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Key not found\n"})
	}
}

func FlushallHandler(store *stokage.Store) gin.HandlerFunc {
	return func(c *gin.Context) {
		Flushall(c, store)
	}
}

func Flushall(c *gin.Context, store *stokage.Store) {
	ok := store.Flushall()

	if ok {
		c.JSON(http.StatusOK, nil)
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to clear all data\n"})
	}
}
