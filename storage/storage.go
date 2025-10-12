package storage

import "sync"

type Storage struct {
	mu   sync.RWMutex
	urls map[string]string
}

func NewStorage() *Storage {
	return &Storage{
		urls: make(map[string]string),
	}
}

func (s *Storage) Get(code string) (string, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	url, ok := s.urls[code]
	return url, ok
}

func (s *Storage) Save(code, url string){
	s.mu.Lock()
	defer s.mu.Unlock()
	s.urls[code] = url
}