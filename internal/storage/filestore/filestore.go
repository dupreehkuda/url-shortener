package filestore

import (
	"bufio"
	"io"
	"log"
	"math/rand"
	"os"
	"strings"
	"sync"
)

type storage struct {
	mtx    sync.RWMutex
	file   *os.File
	reader *bufio.Reader
	writer *bufio.Writer
}

// FILE_STORAGE_PATH

func New(filepath string) *storage {
	file, err := os.OpenFile(filepath, os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		log.Fatal(err)
	}

	return &storage{
		file:   file,
		reader: bufio.NewReader(file),
		writer: bufio.NewWriter(file),
	}
}

// Get - выдаем нужную ссылку по id, который получили.
func (s *storage) Get(id string) (string, error) {
	s.mtx.RLock()
	defer s.mtx.RUnlock()

	_, err := s.file.Seek(0, io.SeekStart)
	if err != nil {
		return "", err
	}
	for {
		b, _, err := s.reader.ReadLine()
		if err != nil {
			return "", err
		}

		record := strings.Split(string(b), ",")
		if record[0] == id {
			return record[1], nil
		}
	}
}

// Create - записываем в хранилище.
func (s *storage) Create(link string) (string, error) {
	s.mtx.Lock()
	s.file.Seek(0, io.SeekEnd)
	codedName := randSymbols(7)

	record := []string{codedName, link}

	_, err := s.writer.WriteString(strings.Join(record, ",") + "\n")
	if err != nil {
		return "", err
	}
	s.writer.Flush()
	s.mtx.Unlock()
	return codedName, nil
}

func randSymbols(n int) string {
	const letterBytes = "1234567890abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}

	return string(b)
}
