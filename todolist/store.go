package todolist

//Store Describes what is need to build a Todo store
type Store interface {
	// Store manipulation
	Initialize() error
	Load() error
	Save() error
	// CRUD (ish) functions
	Size() int
	// CREATE
	Add(todo *Todo) error
	// RETRIEVE a single ID
	FetchById(id int) (*Todo, error)
	// RETRIEVE all the todos
	FetchAll() ([]*Todo, error)
	// DELETE
	Delete(id int) error
}
