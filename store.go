package redis

type store struct {
	data map[string]any
}

func OpenStore() *store {
	return &store{
		data: make(map[string]any),
	}
}

func (s *store) Set(k string, v any) {
	s.data[k] = v
}

func (s *store) Get(k string) (any, error) {
	v, ok := s.data[k]
	if !ok {
		return nil, ErrNotFound
	}
	return v, nil
}
