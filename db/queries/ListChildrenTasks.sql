-- name: ListChildrenTasks :many
SELECT
    task_id,
    parent_task_id,
    task_order,
    name,
    description,
    start_time,
    end_time,
    updated_ts,
    updated_user
FROM task
WHERE parent_task_id = $1
ORDER BY task_order;
