package store

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

type JsonStore struct {
	Store

	memory   map[int]*Todo
	filePath string
	currId   int
}

type jsonFileFormat struct {
	CurrId int              `json:"current_id"`
	Data   map[string]*Todo `json:"data"`
}

func NewJsonStore(filePath string) (*JsonStore, error) {
	file, err := os.Open(filePath)

	if err != nil {
		if strings.Contains(err.Error(), "cannot find the file specified") {
			store := &JsonStore{
				memory:   map[int]*Todo{},
				filePath: filePath,
				currId:   0,
			}
			err := store.save()

			if err != nil {
				return nil, err
			}

			return store, nil

		} else {
			return nil, err
		}
	} else {
		defer func(file *os.File) {
			err := file.Close()
			if err != nil {
				log.Fatal(err)
			}
		}(file)
	}

	// read our opened jsonFile as a byte array.
	byteValue, _ := io.ReadAll(file)

	// we initialize our Users array
	var data jsonFileFormat

	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'users' which we defined above
	err = json.Unmarshal(byteValue, &data)

	if err != nil {
		return nil, err
	}

	store := &JsonStore{
		memory:   map[int]*Todo{},
		filePath: filePath,
		currId:   data.CurrId,
	}

	for id, todo := range data.Data {
		conv, _ := strconv.Atoi(id)
		store.memory[conv] = todo
	}

	return store, nil
}

func (store *JsonStore) save() error {
	jsonData := &jsonFileFormat{
		CurrId: store.currId,
		Data:   map[string]*Todo{},
	}

	for id, todo := range store.memory {
		jsonData.Data[strconv.Itoa(id)] = todo
	}

	jsonString, _ := json.Marshal(jsonData)
	return os.WriteFile(store.filePath, jsonString, os.ModePerm)
}

func (store *JsonStore) GetAll() []*Todo {
	var todos = make([]*Todo, 0)

	for _, todo := range store.memory {
		todos = append(todos, todo)
	}

	return todos
}

func (store *JsonStore) GetById(id int) *Todo {
	return store.memory[id]
}

func (store *JsonStore) GetByName(name string) *Todo {
	for _, todo := range store.memory {
		if todo.Name == name {
			return todo
		}
	}

	return nil
}

func (store *JsonStore) Create(name string, description string) (*Todo, error) {
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

	err := store.save()
	if err != nil {
		return nil, err
	}

	return store.memory[store.currId], nil
}

func (store *JsonStore) Delete(id int) error {
	delete(store.memory, id)

	return store.save()
}
