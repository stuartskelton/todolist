package todolist

import (
	"fmt"
)

//MemoryStore holds the todos in memory
type MemoryStore struct {
	Todos []*Todo
}

//NewMemoryStore Creates a new Memory Store
func NewMemoryStore() *MemoryStore {
	return &MemoryStore{}
}

func (m *MemoryStore) Initialize() error {
	return nil
}

func (m *MemoryStore) Load() error {
	fmt.Printf("MEM! ---- %+v", m.Todos)

	return nil
}

func (m *MemoryStore) Save() error {
	return nil
}

//Add adds a todo to the memory store
func (m *MemoryStore) Add(todo *Todo) error {
	todo.Id = m.NextID()
	m.Todos = append(m.Todos, todo)
	return nil
}

//Fetch fetches a todo based on its id
func (m *MemoryStore) FetchById(id int) (*Todo, error) {
	for _, todo := range m.Todos {
		if todo.Id == id {
			return todo, nil
		}
	}
	return nil, fmt.Errorf("Could not find the todo")
}

//FetchAll Fetches all the todos for this store
func (m *MemoryStore) FetchAll() ([]*Todo, error) {
	return m.Todos, nil
}

//Delete Deletes a todo based on its id
func (m *MemoryStore) Delete(id int) error {
	index := m.indexOf(id)
	m.Todos = append(m.Todos[:index], m.Todos[index+1:]...)
	return nil
}

//Size Returns the current size of the todolist
func (m *MemoryStore) Size() int {
	return len(m.Todos)
}

// useful funcs

//MaxID current max id of the todos in memory
func (m *MemoryStore) MaxID() int {
	maxID := 0
	for _, todo := range m.Todos {
		if todo.Id > maxID {
			maxID = todo.Id
		}
	}
	return maxID
}

//NextID returns the next id
func (m *MemoryStore) NextID() int {
	var found bool
	maxID := m.MaxID()
	for i := 1; i <= maxID; i++ {
		found = false
		for _, todo := range m.Todos {
			if todo.Id == i {
				found = true
				break
			}
		}
		if !found {
			return i
		}
	}
	return maxID + 1
}

func (m *MemoryStore) indexOf(ID int) int {
	for i, todo := range m.Todos {
		if todo.Id == ID {
			return i
		}
	}
	return -1
}
