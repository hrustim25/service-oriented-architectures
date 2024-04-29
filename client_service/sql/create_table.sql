CREATE TABLE IF NOT EXISTS users (
    user_id SERIAL PRIMARY KEY,
    login VARCHAR(25) NOT NULL UNIQUE,
    pwd_hash VARCHAR(100) NOT NULL,
    name VARCHAR(100) DEFAULT '',
    surname VARCHAR(100) DEFAULT '',
    birthdate VARCHAR(20) DEFAULT '',
    email VARCHAR(100) DEFAULT '',
    phone_number VARCHAR(20) DEFAULT ''
);
