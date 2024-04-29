SELECT task_id, name, description, deadline_date, creation_date, completion_status
FROM tasks
WHERE task_id=$1
