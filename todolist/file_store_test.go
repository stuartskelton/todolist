package todolist

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFileStore(t *testing.T) {
	assert := assert.New(t)
	store := NewFileStoreWithLocation("todos.json")
	fmt.Printf("%v+", store)
	assert.Equal(1, 1)
	// todo, _ := store.FetchById(1)
	// assert.Equal(todo.Subject, "this is the first subject", "")
}
