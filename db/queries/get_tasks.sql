-- name: GetTopLevelTasks :many
SELECT
    t.task_id,
    t.name,
    t.description,
    ts.status_name AS status,
    t.start_time,
    t.end_time,
    t.updated_user
FROM task t
         JOIN task_status ts ON t.status_id = ts.status_id
WHERE t.parent_task_id IS NULL
ORDER BY t.updated_ts DESC;
