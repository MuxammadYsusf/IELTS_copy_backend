create table l_questions (
    id serial primary key,
    contenrt_id int,
    order_id int,
    body text,
    true_answer text,
    FOREIGN KEY (test_id) REFERENCES tests(id) ON DELETE CASCADE,
    FOREIGN KEY (contenrt_id) REFERENCES l_contents(id) ON DELETE CASCADE
);

create table l_contents(
    id serial primary key,
    body text,
    section_id int,
    test_id int,
    FOREIGN KEY (test_id) REFERENCES tests(id) ON DELETE CASCADE
)