-- +goose Up
CREATE TABLE student_credentials(
    student_id INT PRIMARY KEY,
    password_hash TEXT NOT NULL
);

-- +goose Down
DROP TABLE student_credentials;