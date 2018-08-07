package todolist

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/user"
)

//FileStore Holds the state for a filestore
type FileStore struct {
	MemoryStore
	FileLocation string
	Loaded       bool
}

//NewFileStore Create a new File store
func NewFileStore() *FileStore {
	return &FileStore{FileLocation: "", Loaded: false}
}

//NewFileStoreWithLocation Create a new File store with the given location
func NewFileStoreWithLocation(location string) *FileStore {
	return &FileStore{FileLocation: location, Loaded: false}
}

//Initialize Initialize json file.
func (f *FileStore) Initialize() error {
	if f.FileLocation == "" {
		f.FileLocation = getLocation()
	}

	_, err := ioutil.ReadFile(f.FileLocation)
	if err == nil {
		return fmt.Errorf("It looks like a .todos.json file already exists!  Doing nothing")
	}
	if err := ioutil.WriteFile(f.FileLocation, []byte("[]"), 0644); err != nil {
		return fmt.Errorf("Error writing json file %s", err)
	}
	fmt.Println("Todo repo initialized.")
	return nil
}

//Load Load the internal Todolist from a JSON file
func (f *FileStore) Load() error {
	if f.FileLocation == "" {
		f.FileLocation = ".todos.json"
	}
	fmt.Printf("LOCARION    %v\n", f.FileLocation)
	data, err := ioutil.ReadFile(f.FileLocation)
	if err != nil {
		fmt.Println("No todo file found!")
		fmt.Println("Initialize a new todo repo by running 'todo init'")
		return err
	}

	var todos []*Todo
	jerr := json.Unmarshal(data, &todos)
	if jerr != nil {
		fmt.Println("Error reading json data", jerr)
		return jerr
	}
	f.Loaded = true
	f.Todos = todos
	fmt.Printf("FILE ---- %+v", todos)
	return nil
}

//Save Save the internal Todolist to the Json location
func (f *FileStore) Save() error {
	data, _ := json.Marshal(f.Todos)
	if err := ioutil.WriteFile(f.FileLocation, []byte(data), 0644); err != nil {
		fmt.Println("Error writing json file", err)
		return err
	}
	return nil
}

//RemoveFile Removes the file
func (f *FileStore) RemoveFile() error {
	error := os.Remove(f.FileLocation)
	if error != nil {
		return error
	}
	return nil
}

func getLocation() string {
	localrepo := ".todos.json"
	usr, _ := user.Current()
	homerepo := fmt.Sprintf("%s/.todos.json", usr.HomeDir)
	_, ferr := os.Stat(localrepo)

	if ferr == nil {
		return localrepo
	}
	return homerepo

}
