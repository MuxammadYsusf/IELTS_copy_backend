CREATE TABLE if not exists users(
    id SERIAL PRIMARY KEY,
    name text,
    phonenumber text,
    password text,
    FOREIGN KEY (id) REFERENCES result(user_id) ON DELETE CASCADE
);