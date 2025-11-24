create table r_questions (
    id serial primary key,
    question text,
    pessage_id int,
    test_id int,
    order_id int,
    body text,
    FOREIGN KEY(test_id) REFERENCES test(id) ON DELETE CASCADE
);

create table r_contents(
    id serial primary key,
    body text,
    pessage_id int,
    test_id int,
    FOREIGN KEY (test_id) REFERENCES test(id) ON DELETE CASCADE
);