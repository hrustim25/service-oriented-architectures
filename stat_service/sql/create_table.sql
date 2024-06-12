CREATE TABLE IF NOT EXISTS events (
    event_id SERIAL PRIMARY KEY,
    task_id BIGINT NOT NULL,
    task_author_id BIGINT NOT NULL,
    user_id BIGINT NOT NULL,
    event_type INTEGER NOT NULL,
    event_date VARCHAR(20) DEFAULT '',
    UNIQUE (task_id, user_id, event_type)
);
