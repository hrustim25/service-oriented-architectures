SELECT task_id, name, description, deadline_date, creation_date, completion_status
FROM tasks
WHERE author_id=$1
LIMIT $2
OFFSET $3
