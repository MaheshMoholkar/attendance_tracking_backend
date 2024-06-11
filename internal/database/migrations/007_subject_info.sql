-- +goose Up
CREATE TABLE subject_info (
    subject_id SERIAL PRIMARY KEY,
    subjectName TEXT NOT NULL,
    class_id INT NOT NULL REFERENCES class_info(class_id) ON DELETE CASCADE
);

-- +goose Down
DROP TABLE subject_info;