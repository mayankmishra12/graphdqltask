package graph

import "github.com/mmishra12/gqlgen-todos/internal/db"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	DB *db.Queries
}

func (r *Resolver) Query() QueryResolver {
	//TODO implement me
	panic("implement me")
}
