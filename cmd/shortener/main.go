package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
)

var shortened = make(map[int]string)

func trimmer(path string) int {
	id := strings.TrimPrefix(path, "/")
	i, err := strconv.Atoi(id)

	if err != nil {
		return 0
	}

	return i
}

func apiResponse(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		i := trimmer(r.URL.Path)

		if len(shortened) == 0 || i == 0 || shortened[i] == "" {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Can't find url requested"))
			return
		}

		w.Header().Set("Location", shortened[i])
		w.WriteHeader(http.StatusTemporaryRedirect)

	case http.MethodPost:
		b, err := io.ReadAll(r.Body)
		if err != nil || len(b) == 0 {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		shortened[len(shortened)+1] = string(b)
		string := fmt.Sprintf("http://localhost:8080/%v", len(shortened))
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(string))

	default:
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Can't find method requested"))
	}
}

func main() {
	http.HandleFunc("/", apiResponse)
	if err := http.ListenAndServe("localhost:8080", nil); err != nil {
		log.Fatal(err)
	}
}
