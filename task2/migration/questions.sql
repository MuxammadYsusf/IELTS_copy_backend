CREATE TABLE if not exists questions(
    id SERIAL PRIMARY KEY,
    name text,
    a text,
    b text,
    true_answer text,
    subject_id int,
    grade_id int,
    FOREIGN KEY (id) REFERENCES result(question_id) ON DELETE CASCADE
);