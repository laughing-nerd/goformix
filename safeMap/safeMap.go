package safemap

import "sync"

type SafeMap struct {
	mu   sync.Mutex
	data map[string]interface{}
}

func New() SafeMap {
	return SafeMap{
		data: make(map[string]interface{}),
	}
}

func (s *SafeMap) Set(key string, value interface{}) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.data[key] = value
}

func (s *SafeMap) Get(key string) (interface{}, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	value, ok := s.data[key]
	return value, ok
}

func (s *SafeMap) GetMap() map[string]interface{} {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.data
}
