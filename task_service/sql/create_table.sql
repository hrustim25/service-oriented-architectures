CREATE TABLE IF NOT EXISTS tasks (
    task_id SERIAL PRIMARY KEY,
    author_id BIGINT NOT NULL,
    name VARCHAR(100) DEFAULT '',
    description VARCHAR(1000) DEFAULT '',
    deadline_date VARCHAR(20) DEFAULT '',
    creation_date VARCHAR(20) DEFAULT '',
    completion_status VARCHAR(20) DEFAULT 'NOT STARTED'
);
