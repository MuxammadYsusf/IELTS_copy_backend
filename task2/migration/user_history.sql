CREATE TABLE if not exists attempts(
    id SERIAL PRIMARY KEY,
    created_at timestamp,
    FOREIGN KEY (id) REFERENCES result(attempt_id) ON DELETE CASCADE
);