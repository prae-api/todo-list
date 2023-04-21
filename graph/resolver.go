package graph

import "github.com/prae-api/todo-api/graph/model"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	TodoStore map[int]model.Todo
}
