package testhelpers

type DummyStore struct {
	db map[string]string
}

func (s *DummyStore) Open(_ string) error {
	return nil
}

func (s *DummyStore) Close() error {
	return nil
}

func (s *DummyStore) Get(key []byte) ([]byte, error) {
	data, ok := s.db[string(key)]
	if !ok {
		return nil, nil
	}
	return []byte(data), nil
}

func (s *DummyStore) Put(key []byte, value []byte) error {
	s.db[string(key)] = string(value)
	return nil
}

func (s *DummyStore) Status() map[string]string {
	return map[string]string{
		"driver": "dummy",
	}
}

func (s *DummyStore) String() string {
	return "DummyStore"
}
