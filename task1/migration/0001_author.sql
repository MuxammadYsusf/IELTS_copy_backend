CREATE TABLE if not exists authors(
    author_id INT PRIMARY KEY,
    author_name TEXT,
    password TEXT,
    FOREIGN KEY (Author_id) REFERENCES authors(author_id) ON DELETE CASCADE
);