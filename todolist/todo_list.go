package todolist

import (
	"sort"
)

type TodoList struct {
	Store Store
}

//NewTodoList Create a new instance of todolist with the provided store.
func NewTodoList(storage Store) *TodoList {
	return &TodoList{Store: storage}
}

func (t *TodoList) Load(todos []*Todo) {
	t.Store.Load()
}

func (t *TodoList) FindById(id int) *Todo {
	todo, err := t.Store.FetchById(id)
	if err != nil {
		panic(err)
	}
	return todo
}

func (t *TodoList) Add(todo *Todo) {
	t.Store.Add(todo)
}

func (t *TodoList) Delete(ids ...int) {
	for _, id := range ids {
		t.Store.Delete(id)
	}
}

func (t *TodoList) Complete(ids ...int) {
	for _, id := range ids {
		todo, _ := t.Store.FetchById(id)
		if todo == nil {
			continue
		}
		todo.Complete()
	}
}

func (t *TodoList) Uncomplete(ids ...int) {
	for _, id := range ids {
		todo, _ := t.Store.FetchById(id)
		if todo == nil {
			continue
		}
		todo.Uncomplete()
	}
}

func (t *TodoList) Archive(ids ...int) {
	for _, id := range ids {
		todo, _ := t.Store.FetchById(id)
		if todo == nil {
			continue
		}
		todo.Archive()
	}
}

func (t *TodoList) Unarchive(ids ...int) {
	for _, id := range ids {
		todo, _ := t.Store.FetchById(id)
		if todo == nil {
			continue
		}
		todo.Unarchive()
	}
}

func (t *TodoList) Prioritize(ids ...int) {
	for _, id := range ids {
		todo, _ := t.Store.FetchById(id)
		if todo == nil {
			continue
		}
		todo.Prioritize()
	}
}

func (t *TodoList) Unprioritize(ids ...int) {
	for _, id := range ids {
		todo, _ := t.Store.FetchById(id)
		if todo == nil {
			continue
		}
		todo.Unprioritize()
	}
}

type ByDate []*Todo

func (a ByDate) Len() int      { return len(a) }
func (a ByDate) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a ByDate) Less(i, j int) bool {
	t1Due := a[i].CalculateDueTime()
	t2Due := a[j].CalculateDueTime()
	return t1Due.Before(t2Due)
}

func (t *TodoList) Todos() []*Todo {
	todos, _ := t.Store.FetchAll()
	sort.Sort(ByDate(todos))
	return todos
}

func (t *TodoList) GarbageCollect() {
	var toDelete []*Todo
	todos, _ := t.Store.FetchAll()
	for _, todo := range todos {
		if todo.Archived {
			toDelete = append(toDelete, todo)
		}
	}
	for _, todo := range toDelete {
		t.Delete(todo.Id)
	}
}
