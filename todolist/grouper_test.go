package todolist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGroupByContext(t *testing.T) {
	assert := assert.New(t)

	list := SetUpMemoryTodoList()
	todos, _ := list.Store.FetchAll()

	grouper := &Grouper{}
	grouped := grouper.GroupByContext(todos)

	assert.Equal(2, len(grouped.Groups["root"]), "")
	assert.Equal(1, len(grouped.Groups["more"]), "")
}

func TestGroupByProject(t *testing.T) {
	assert := assert.New(t)

	list := SetUpMemoryTodoList()
	todos, _ := list.Store.FetchAll()

	grouper := &Grouper{}
	grouped := grouper.GroupByProject(todos)

	assert.Equal(2, len(grouped.Groups["test1"]), "")
}
