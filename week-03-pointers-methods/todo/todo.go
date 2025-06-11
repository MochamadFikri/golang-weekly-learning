package todo

import "fmt"

type Todo struct {
	ID        int
	Task      string
	Completed bool
}

type TodoList struct {
	todos  []*Todo
	nextID int
}

func NewTodoList() *TodoList {
	return &TodoList{
		todos:  []*Todo{},
		nextID: 1,
	}
}

func (tl *TodoList) Add(task string) {
	todo := &Todo{
		ID:   tl.nextID,
		Task: task,
	}
	tl.todos = append(tl.todos, todo)
	tl.nextID++
}

func (tl *TodoList) Complete(id int) error {
	for _, todo := range tl.todos {
		if todo.ID == id {
			todo.Completed = true
			return nil
		}
	}
	return fmt.Errorf("Todo dengan ID %d tidak ditemukan", id)
}

func (tl *TodoList) Delete(id int) error {
	for i, todo := range tl.todos {
		if todo.ID == id {
			tl.todos = append(tl.todos[:i], tl.todos[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("Todo dengan ID %d tidak ditemukan", id)
}

func (tl TodoList) ShowAll() {
	if len(tl.todos) == 0 {
		fmt.Println("Tidak ada todo.")
		return
	}
	for _, todo := range tl.todos {
		status := "❌"
		if todo.Completed {
			status = "✅"
		}
		fmt.Printf("[%s] %d: %s\n", status, todo.ID, todo.Task)
	}
}
