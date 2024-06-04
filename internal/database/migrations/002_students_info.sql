-- +goose Up
CREATE TABLE student_info (
    id SERIAL PRIMARY KEY,
    firstName TEXT NOT NULL,
    lastName TEXT NOT NULL,
    rollno INT NOT NULL,
    email TEXT NOT NULL,
    className TEXT NOT NULL,
    division TEXT NOT NULL,
    year INT NOT NULL,
    student_id INT NOT NULL REFERENCES student_credentials(student_id) ON DELETE CASCADE,
    CONSTRAINT unique_student_id UNIQUE (student_id)
);

-- +goose Down
DROP TABLE student_info;
