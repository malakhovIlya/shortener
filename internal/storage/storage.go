package storage

type Storage interface {
	Save(code, longURL string) error
	Get(code string) (string, error)
}

type InMemoryStorage struct {
	Data map[string]string
}
