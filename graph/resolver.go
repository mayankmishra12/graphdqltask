package graph

import (
	"context"
	"github.com/google/uuid"
	"github.com/mmishra12/gqlgen-todos/graph/model"
	"github.com/mmishra12/gqlgen-todos/internal/db"
)

// // This file will not be regenerated automatically.
// //
// // It serves as dependency injection for your app, add any dependencies you require here.
type Resolver struct {
	DB *db.Queries
}

//
//func (r *Resolver) Query() QueryResolver {
//	//TODO implement me
//	panic("implement me")
//}

func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

type queryResolver struct{ *Resolver }

func (q *queryResolver) Task(ctx context.Context, id uuid.UUID) (*model.Task, error) {
	return q.Resolver.Task(ctx, id)
}

func (q *queryResolver) RootTasks(ctx context.Context, first *int32, after *string, last *int32, before *string, sort []*model.TaskSort) (*model.TaskConnection, error) {
	return q.Resolver.RootTasks(ctx, first, after, last, before, sort)
}
