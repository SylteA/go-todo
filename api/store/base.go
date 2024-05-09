package store

type Todo struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type Store interface {
	GetAll() []*Todo
	GetById(id int) *Todo
	GetByName(name string) *Todo
	Create(name string, description string) (*Todo, error)
	Delete(id int) error
}

type StorageType = int

const (
	InMemory StorageType = iota
	JSON     StorageType = iota
)
