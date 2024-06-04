-- +goose Up
CREATE TABLE attendance_info (
    attendance_id SERIAL PRIMARY KEY,
    student_id INT NOT NULL REFERENCES student_info(student_id) ON DELETE CASCADE,
    date DATE NOT NULL,
    status BOOLEAN NOT NULL,
    class_id INT NOT NULL REFERENCES class_info(class_id) ON DELETE CASCADE,
    division_id INT NOT NULL REFERENCES division_info(division_id) ON DELETE CASCADE
);

-- +goose Down
DROP TABLE attendance_info;
