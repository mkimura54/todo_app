package main

import (
	"math"
	"time"
)

type TodoList struct {
	Todos []Todo
}

func (todos *TodoList) AddTodo(content string) {
	todo := Todo{
		No:      todos.GetNextNo(),
		Content: content,
		AddDate: time.Now(),
	}
	todos.Todos = append(todos.Todos, todo)
}

func (todos *TodoList) DeleteTodo(no int) {
	var newTodos []Todo
	for _, todo := range todos.Todos {
		if todo.No != no {
			newTodos = append(newTodos, todo)
		}
		todos.Todos = newTodos
	}
}

func (todos *TodoList) GetNextNo() int {
	var maxNo int
	for _, todo := range todos.Todos {
		maxNo = int(math.Max(float64(maxNo), float64(todo.No)))
	}
	return maxNo + 1
}

type Todo struct {
	No      int
	Content string
	AddDate time.Time
}
