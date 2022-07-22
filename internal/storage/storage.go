package storage

import (
	"errors"
	"math/rand"
	"sync"
)

type storage struct {
	mtx       sync.RWMutex
	shortened map[string]string
}

func New() *storage {
	return &storage{shortened: map[string]string{}}
}

// Get - выдаем нужную ссылку по id, который получили.
func (s *storage) Get(id string) (string, error) {
	if len(s.shortened) == 0 || len(id) == 0 || s.shortened[id] == "" {
		return "", errors.New("can't find url requested")
	}

	s.mtx.RLock()
	defer s.mtx.RUnlock()
	return s.shortened[id], nil
}

// Create - проверям есть ли ссылка и записываем в хранилище
func (s *storage) Create(link string) (string, error) {
	codedName := randSymbols(5)

	if _, ok := s.shortened[codedName]; ok {
		return "", errors.New("this codename already exists")
	}

	key, exist := checkForValue(link, s.shortened)
	if exist {
		return key, nil
	}

	s.mtx.Lock()
	s.shortened[codedName] = link
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

func checkForValue(userValue string, students map[string]string) (string, bool) {
	for key, value := range students {
		if value == userValue {
			return key, true
		}
	}
	return "", false
}
