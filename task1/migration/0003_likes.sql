
CREATE TABLE if not exists likes(
Id INT PRIMARY KEY,
Author_id INT,
FOREIGN KEY (Author_id) REFERENCES authors(author_id) ON DELETE CASCADE
FOREIGN KEY (id) REFERENCES smedia(Id) ON DELETE CASCADE
);
