SELECT COUNT(*)
FROM events
WHERE task_author_id=$1 AND event_type=$2
