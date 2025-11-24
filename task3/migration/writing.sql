create table writing (
    id serial primary key,
    question text,
    test_id int,
    task_id int,
    FOREIGN KEY(test_id) REFERENCES test(id) ON DELETE CASCADE
);