package store

import (
	"fmt"
)

type InMemoryStore struct {
	Store
	memory map[int]*Todo

	currId int
}

func NewInMemoryStore(startingId ...int) (*InMemoryStore, error) {
	store := &InMemoryStore{
		memory: map[int]*Todo{},
		currId: 0,
	}

	if len(startingId) > 0 {
		store.currId = startingId[0]
	}

	if store.currId < 0 {
		return nil, fmt.Errorf("starting id must be greater than 0")
	}

	return store, nil
}

func (store *InMemoryStore) GetAll() []*Todo {
	var todos []*Todo

	for _, todo := range store.memory {
		todos = append(todos, todo)
	}

	return todos
}

func (store *InMemoryStore) GetById(id int) *Todo {
	return store.memory[id]
}

func (store *InMemoryStore) GetByName(name string) *Todo {
	for _, todo := range store.memory {
		if todo.Name == name {
			return todo
		}
	}

	return nil
}

func (store *InMemoryStore) Create(name string, description string) (*Todo, error) {
	existing := store.GetByName(name)

	if existing != nil {
		return nil, fmt.Errorf("todo with name %s already exists", name)
	}

	store.currId++
	store.memory[store.currId] = &Todo{
		ID:          store.currId,
		Name:        name,
		Description: description,
	}

	return store.memory[store.currId], nil
}

func (store *InMemoryStore) Delete(id int) error {
	delete(store.memory, id)
	return nil
}
