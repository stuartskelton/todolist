package todolist

import "testing"

func TestNewTodo(t *testing.T) {
	todo := NewTodo()

	if todo.Completed || todo.Archived || todo.CompletedDate != "" {
		t.Error("Completed should be false for new todos")
	}
}

func TestValidity(t *testing.T) {
	todo := &Todo{Subject: "test"}
	if !todo.Valid() {
		t.Error("Expected valid todo to be valid")
	}

	invalidTodo := &Todo{Subject: ""}
	if invalidTodo.Valid() {
		t.Error("Invalid todo is being reported as valid")
	}
}

// A util that sets up a fixture memory todolist

func SetUpMemoryTodoList() *TodoList {
	list := &TodoList{Store: NewMemoryStore()}
	todo1 := NewTodo()
	todo1.Subject = "hello"
	todo1.Projects = []string{"test1", "test2"}
	todo1.Contexts = []string{"root", "more"}
	list.Add(todo1)

	todo2 := NewTodo()
	todo2.Subject = "hello"
	todo2.Projects = []string{"test1", "test3"}
	todo2.Contexts = []string{"root", "boot"}
	list.Add(todo2)

	todo3 := NewTodo()
	todo3.Subject = "hello"
	todo3.Projects = []string{"test1", "test3"}
	todo3.Contexts = []string{"root", "boot"}
	todo3.Prioritize()
	list.Add(todo3)

	todo4 := NewTodo()
	todo4.Subject = "hello"
	todo4.Projects = []string{"test1", "test3"}
	todo4.Contexts = []string{"root", "boot"}
	todo4.Archive()
	list.Add(todo4)

	return list
}
