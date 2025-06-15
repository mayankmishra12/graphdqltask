-- name: GetTaskByID :one
SELECT
    t.task_id,
    t.parent_task_id,
    t.task_order,
    t.name,
    t.description,
    t.start_time,
    t.end_time,
    t.updated_ts,
    t.updated_user,
    t.last_modified_process,
    t.last_modified_app,
    t.last_request_id,
    t.last_action_id,

    ts.status_id,
    ts.status_name,

    parent.task_id AS parent_id,
    parent.name AS parent_name,

    tt.task_type_id,
    tt.type,

    rj.redjade_task_id,
    rj.redjade_link,

    ss.survey_id

FROM task t
         LEFT JOIN task_status ts ON t.status_id = ts.status_id
         LEFT JOIN task parent ON t.parent_task_id = parent.task_id
         LEFT JOIN task_type tt ON t.task_id = tt.task_id
         LEFT JOIN redjade_task rj ON tt.task_type_id = rj.task_type_id
         LEFT JOIN sams_survey_task ss ON tt.task_type_id = ss.task_type_id
WHERE t.task_id = $1;
