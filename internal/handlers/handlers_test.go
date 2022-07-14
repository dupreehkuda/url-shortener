package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPostShorten(t *testing.T) {
	rPath := "/1"
	router := gin.Default()
	router.GET(rPath, GetShortened(map[int]string{1: "https://youtube.com/"}))
	req, _ := http.NewRequest("GET", rPath, nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	t.Logf("status: %d", w.Code)
	t.Logf("response: %s", w.Body.String())
}
