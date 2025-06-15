-- name: ListRootTasks :many
SELECT
    t.task_id,
    t.task_order,
    t.name,
    t.description,
    t.start_time,
    t.end_time,
    t.updated_ts,
    t.updated_user,
    ts.status_id,
    ts.status_name
FROM task t
         JOIN task_status ts ON t.status_id = ts.status_id
WHERE t.parent_task_id IS NULL
  AND ($1::timestamp IS NULL OR t.updated_ts > $1)
ORDER BY t.updated_ts ASC
    LIMIT $2;
