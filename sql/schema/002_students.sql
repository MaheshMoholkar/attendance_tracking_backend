-- +goose Up
CREATE TABLE students (
    id SERIAL PRIMARY KEY,
    firstName TEXT NOT NULL,
    lastName TEXT NOT NULL,
    rollno INT NOT NULL,
    email TEXT,
    className TEXT NOT NULL,
    division TEXT NOT NULL,
    year INT NOT NULL,
    student_id INT NOT NULL REFERENCES student_credentials(student_id) ON DELETE CASCADE
);

-- +goose Down
DROP TABLE students;