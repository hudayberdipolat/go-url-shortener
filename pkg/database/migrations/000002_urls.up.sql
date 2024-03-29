CREATE TABLE IF NOT EXISTS urls(
    id SERIAL PRIMARY KEY,
    url_name VARCHAR(255),
    short_url VARCHAR(255) NOT NULL UNIQUE,
    long_url VARCHAR(255) NOT NULL,
    visited_count INTEGER DEFAULT 0,
    user_id INTEGER,
    created_at TIMESTAMP,
    CONSTRAINT fk_user FOREIGN KEY(user_id) REFERENCES users(id)
);