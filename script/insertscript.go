package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

var dsn = "postgres://postgres:mysecretpass@localhost:5432/task?sslmode=disable"

func main() {
	ctx := context.Background()

	dbpool, err := pgxpool.New(ctx, dsn)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}
	defer dbpool.Close()

	fmt.Println("Connected to database.")

	if err := initDB(ctx, dbpool); err != nil {
		log.Fatalf("DB initialization failed: %v", err)
	}

	// Insert 10 tasks
	for i := 1; i <= 10; i++ {
		taskID := uuid.New()
		name := fmt.Sprintf("Task #%02d", i)
		description := fmt.Sprintf("Description for task %02d", i)
		taskOrder := i

		// Use real timestamps with date and time:
		startTime := time.Now().Add(time.Duration(i) * time.Hour)
		endTime := startTime.Add(2 * time.Hour)
		updatedTS := time.Now()

		statusID := (i % 3) + 1

		query := `
			INSERT INTO task (
				task_id,
				parent_task_id,
				task_order,
				name,
				description,
				start_time,
				end_time,
				status_id,
				updated_ts,
				updated_user,
				last_modified_process,
				last_modified_app,
				last_request_id,
				last_action_id
			) VALUES ($1, NULL, $2, $3, $4, $5, $6, $7, $8, 'admin', 'processX', 'appX', $9, $10)
		`

		_, err := dbpool.Exec(ctx, query,
			taskID,
			taskOrder,
			name,
			description,
			startTime,
			endTime,
			statusID,
			updatedTS,
			fmt.Sprintf("req-%03d", i),
			fmt.Sprintf("act-%03d", i),
		)

		if err != nil {
			log.Printf("Failed to insert task %d: %v\n", i, err)
		} else {
			log.Printf("Inserted task %d (%s)\n", i, taskID)
		}
	}
}

func initDB(ctx context.Context, dbpool *pgxpool.Pool) error {
	statements := []string{
		`CREATE TABLE IF NOT EXISTS task_status (
			status_id INT PRIMARY KEY,
			status_name TEXT NOT NULL
		);`,
		`CREATE TABLE IF NOT EXISTS task (
			task_id UUID PRIMARY KEY,
			parent_task_id UUID,
			task_order INT NOT NULL,
			name TEXT NOT NULL,
			description TEXT,
			start_time TIMESTAMP,
			end_time TIMESTAMP,
			status_id INT NOT NULL REFERENCES task_status(status_id),
			updated_ts TIMESTAMP NOT NULL,
			updated_user TEXT NOT NULL,
			last_modified_process TEXT,
			last_modified_app TEXT,
			last_request_id TEXT NOT NULL,
			last_action_id TEXT
		);`,
		`INSERT INTO task_status (status_id, status_name) VALUES
			(1, 'Pending'),
			(2, 'In Progress'),
			(3, 'Completed')
		ON CONFLICT (status_id) DO NOTHING;`,
	}

	for _, stmt := range statements {
		_, err := dbpool.Exec(ctx, stmt)
		if err != nil {
			return fmt.Errorf("init failed: %w", err)
		}
	}

	fmt.Println("Database initialized.")
	return nil
}
