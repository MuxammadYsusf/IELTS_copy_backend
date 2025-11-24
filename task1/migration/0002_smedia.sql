CREATE TABLE if not exists smedia(
    Id INT PRIMARY KEY,
    Title TEXT,
    Description TEXT,
    Author_id INT,
    Date BIGINT,
    FOREIGN KEY(Author_id) REFERENCES authors(Author_id) ON DELETE CASCADE

);
