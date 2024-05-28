SELECT task_id, task_author_id
FROM (SELECT task_id, task_author_id, COUNT(task_id)
        FROM events
        WHERE event_type=$1
        GROUP BY task_id, task_author_id) AS top_tasks (task_id, task_author_id, event_count)
ORDER BY event_count DESC
LIMIT 5
