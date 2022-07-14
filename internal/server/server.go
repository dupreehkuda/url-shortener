package server

import (
	"github.com/gin-gonic/gin"
)

func GetRouter() *gin.Engine {
	r := gin.Default()

	// здесь хотел прописать эндпоинты, но не придумал как притянуть storage

	return r
}
