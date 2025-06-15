package graph

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/mmishra12/gqlgen-todos/graph/model"
	"github.com/mmishra12/gqlgen-todos/internal/db"
)

func (r *Resolver) Task(ctx context.Context, id uuid.UUID) (*model.Task, error) {
	// TODO: Use generated GetTaskByID once added via sqlc.
	return nil, fmt.Errorf("not implemented: Task - task")
}

func (r *Resolver) RootTasks(ctx context.Context, first *int32, after *string, last *int32, before *string, sort []*model.TaskSort) (*model.TaskConnection, error) {
	var afterTime pgtype.Timestamp
	if after != nil {
		t, err := time.Parse(time.RFC3339, *after)
		if err != nil {
			return nil, err
		}
		afterTime = pgtype.Timestamp{Time: t, Valid: true}
	} else {
		afterTime = pgtype.Timestamp{Valid: false}
	}

	limit := int32(10) // default
	if first != nil {
		limit = *first
	}

	rows, err := r.DB.ListRootTasks(ctx, db.ListRootTasksParams{
		Column1: afterTime,
		Limit:   limit,
	})
	if err != nil {
		return nil, err
	}

	var edges []*model.TaskEdge
	for _, row := range rows {
		task := &model.Task{
			TaskID: row.TaskID.Bytes,
			Name:   row.Name,
			//Description: row.Description.String,
			TaskOrder:   int32(row.TaskOrder),
			StartTime:   toTimePtr(row.StartTime),
			EndTime:     toTimePtr(row.EndTime),
			UpdatedTs:   row.UpdatedTs.Time,
			UpdatedUser: row.UpdatedUser,
			Status: &model.TaskStatus{
				StatusID:   row.StatusID,
				StatusName: row.StatusName,
			},
		}

		cursor := row.UpdatedTs.Time.Format(time.RFC3339)
		edges = append(edges, &model.TaskEdge{
			Node:   task,
			Cursor: cursor,
		})
	}

	var pageInfo model.PageInfo
	if len(edges) > 0 {
		pageInfo.StartCursor = &edges[0].Cursor
		pageInfo.EndCursor = &edges[len(edges)-1].Cursor
	}
	pageInfo.HasNextPage = len(edges) == int(limit)
	pageInfo.HasPreviousPage = false // not implemented in this direction

	return &model.TaskConnection{
		Edges:    edges,
		PageInfo: &pageInfo,
	}, nil
}

// Converts pgtype.Timestamp to *time.Time
func toTimePtr(ts pgtype.Timestamp) *time.Time {
	if ts.Valid {
		return &ts.Time
	}
	return nil
}
