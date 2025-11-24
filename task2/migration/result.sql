CREATE TABLE if not exists result(
    id SERIAL PRIMARY KEY,
    question_id int,
    is_correct boolean,
    attempt_id int,
    user_id int,
    selected_option text
);