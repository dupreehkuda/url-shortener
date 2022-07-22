package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http"
)

type Storer interface {
	Get(id string) (string, error)
	Create(link string) (string, error)
}

type handlers struct {
	storage Storer
}

func New(storage Storer) *handlers {
	return &handlers{storage: storage}
}

// GetShortened - обрабатываем Get-запрос и переадресуем пользователя
func (h handlers) GetShortened() gin.HandlerFunc {
	fn := func(c *gin.Context) {
		id := c.Param("id")

		ans, err := h.storage.Get(id)
		if err != nil {
			c.Data(http.StatusBadRequest, "text/plain; charset=utf-8", []byte("Can't find url requested"))
			log.Printf("Got: %v, Recieved error: %v", id, err)
			return
		}

		c.Header("Location", ans)
		c.JSON(http.StatusTemporaryRedirect, nil)

		log.Printf("Redirected to %v", ans)
	}

	return gin.HandlerFunc(fn)
}

// PostShorten - обрабатываем Post-запрос и возвращаем ответ
func (h handlers) PostShorten() gin.HandlerFunc {
	fn := func(c *gin.Context) {
		b, err := io.ReadAll(c.Request.Body)

		if err != nil || len(b) == 0 {
			c.JSON(http.StatusBadRequest, nil)
			log.Printf("Recieved error while reading body: %v", err)
			return
		}

		ans, err := h.storage.Create(string(b))
		if err != nil {
			c.JSON(http.StatusBadRequest, nil)
			log.Printf("Recieved error: %v", err)
			return
		}

		responseText := fmt.Sprintf("http://localhost:8080/%v", ans)
		log.Printf("New link responce: %s", responseText)
		c.Data(http.StatusCreated, "text/plain; charset=utf-8", []byte(responseText))
	}

	return gin.HandlerFunc(fn)
}
