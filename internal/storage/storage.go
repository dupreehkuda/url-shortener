package storage

type storage struct {
	Shortened map[int]string
}

func New() *storage {
	return &storage{Shortened: map[int]string{}}
}
