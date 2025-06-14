-- name: GetSubTasksByParent :many
SELECT
    t.task_id,
    t.name,
    t.description,
    ts.status_name AS status,
    t.start_time,
    t.end_time,
    t.updated_user,
    t.task_order,
    tt.type AS task_type,
    rj.redjade_link,
    sst.survey_id
FROM task t
         JOIN task_status ts ON t.status_id = ts.status_id
         LEFT JOIN task_type tt ON t.task_id = tt.task_id
         LEFT JOIN redjade_task rj ON tt.task_type_id = rj.task_type_id
         LEFT JOIN sams_survey_task sst ON tt.task_type_id = sst.task_type_id
WHERE t.parent_task_id = $1
ORDER BY t.task_order;
