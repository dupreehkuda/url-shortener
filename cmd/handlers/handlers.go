package handlers

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
)

func trimmer(path string) int {
	id := strings.TrimPrefix(path, "/")
	i, err := strconv.Atoi(id)

	if err != nil {
		return 0
	}

	return i
}

func APIResponse(storage map[int]string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			i := trimmer(r.URL.Path)

			if len(storage) == 0 || i == 0 || storage[i] == "" {
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte("Can't find url requested"))
				return
			}

			w.Header().Set("Location", storage[i])
			w.WriteHeader(http.StatusTemporaryRedirect)

		case http.MethodPost:
			b, err := io.ReadAll(r.Body)
			if err != nil || len(b) == 0 {
				w.WriteHeader(http.StatusBadRequest)
				return
			}

			storage[len(storage)+1] = string(b)
			string := fmt.Sprintf("http://localhost:8080/%v", len(storage))
			w.WriteHeader(http.StatusCreated)
			w.Write([]byte(string))

		default:
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Can't find method requested"))
		}
	}
}
