create table r_attempts(
    id serial primary key,
    user_id int,
    start_time timestamp,
    end_time timestamp,
    selected_pessage int[],
    order_id int,
    question_id int,
    real_end timestamp,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,

);

create table r_results(
    id serial primary key,
    attempt_id int unique,
    order_id int,
    user_answer text,
    user_id int,
    score int
    FOREIGN KEY (attempt_id) REFERENCES r_attempts(id) ON DELETE CASCADE
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
);