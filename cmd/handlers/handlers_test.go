package handlers

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAPIResponse(t *testing.T) {
	type storage map[int]string

	type want struct {
		contentType string
		statusCode  int
		response    []byte
	}

	tests := []struct {
		name    string
		request string
		storage storage
		data    string
		want    want
	}{
		{
			name:    "post",
			request: "/",
			storage: storage{},
			data:    "https://youtube.com/",
			want: want{
				contentType: "",
				statusCode:  201,
				response:    []byte("http://localhost:8080/1"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			request := httptest.NewRequest(http.MethodPost, tt.request, bytes.NewBuffer([]byte(tt.data)))
			w := httptest.NewRecorder()
			h := http.HandlerFunc(APIResponse(storage{}))
			h.ServeHTTP(w, request)
			result := w.Result()

			defer result.Body.Close()
			
			body, err := ioutil.ReadAll(result.Body)
			if err != nil {
				log.Fatalln(err)
			}

			assert.Equal(t, tt.want.statusCode, result.StatusCode)
			assert.Equal(t, tt.want.contentType, result.Header.Get("Content-Type"))
			assert.Equal(t, tt.want.response, body)
		})
	}
}
