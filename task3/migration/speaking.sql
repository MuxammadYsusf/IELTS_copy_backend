create table speaking (
    id serial primary key,
    question text,
    part_id int,
    test_id int,
    order_id int,
    FOREIGN KEY(test_id) REFERENCES test(id) ON DELETE CASCADE
);