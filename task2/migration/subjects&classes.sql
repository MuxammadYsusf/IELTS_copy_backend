CREATE TABLE if not exists subjects(
    id SERIAL PRIMARY KEY,
    name text,
    FOREIGN KEY (id) REFERENCES questions(subject_id) ON DELETE CASCADE,
    FOREIGN KEY (id) REFERENCES grades(subject_id) ON DELETE CASCADE
);

CREATE TABLE if not exists grades(
    id SERIAL PRIMARY KEY,
    grade_num int,
    subject_id int,
    FOREIGN KEY (id) REFERENCES questions(grade_id) ON DELETE CASCADE
);