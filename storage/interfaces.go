package storage

type DB interface {
	Open(string) error
	Close() error
	Get([]byte) ([]byte, error)
	Put([]byte, []byte) error
	Status() map[string]string
	String() string
}
