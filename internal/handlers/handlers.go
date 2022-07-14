package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"strconv"
)

func GetShortened(storage map[int]string) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return
		}

		if len(storage) == 0 || id == 0 || storage[id] == "" {
			c.Data(http.StatusBadRequest, "text/plain; charset=utf-8", []byte("Can't find url requested"))
			return
		}

		c.Header("Location", storage[id])
		c.JSON(http.StatusTemporaryRedirect, nil)
	}

	return gin.HandlerFunc(fn)
}

func PostShorten(storage map[int]string) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		b, err := io.ReadAll(c.Request.Body)

		if err != nil || len(b) == 0 {
			c.JSON(http.StatusBadRequest, nil)
			return
		}

		storage[len(storage)+1] = string(b)
		string := fmt.Sprintf("http://localhost:8080/%v", len(storage))
		c.Data(http.StatusCreated, "text/plain; charset=utf-8", []byte(string))
	}

	return gin.HandlerFunc(fn)
}
