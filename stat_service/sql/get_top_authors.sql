SELECT task_author_id
FROM (SELECT task_author_id, COUNT(task_author_id)
        FROM events
        WHERE event_type=$1
        GROUP BY task_author_id) AS top_authors (task_author_id, event_count)
ORDER BY event_count DESC
LIMIT 3
