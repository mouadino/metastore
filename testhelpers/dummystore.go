package testhelpers

type InMemoryStore map[string]string

func (s *InMemoryStore) Open(_ string) error {
	return nil
}

func (s *InMemoryStore) Close() error {
	return nil
}

func (s *InMemoryStore) Get(key []byte) ([]byte, error) {
	data, ok := (*s)[string(key)]
	if !ok {
		return nil, nil
	}
	return []byte(data), nil
}

func (s *InMemoryStore) Put(key []byte, value []byte) error {
	(*s)[string(key)] = string(value)
	return nil
}

func (s *InMemoryStore) Status() map[string]string {
	return map[string]string{
		"driver": "dummy",
	}
}

func (s *InMemoryStore) String() string {
	return "InMemoryStore"
}
