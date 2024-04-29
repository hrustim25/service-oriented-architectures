INSERT INTO tasks(author_id, name, description, deadline_date, creation_date)
VALUES($1, $2, $3, $4, $5)
RETURNING task_id
