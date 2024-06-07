-- +goose Up
CREATE TABLE staff_info (
    id SERIAL PRIMARY KEY,
    firstName TEXT NOT NULL,
    lastName TEXT NOT NULL,
    email TEXT NOT NULL,
    class_id INT NOT NULL REFERENCES class_info(class_id),
    staff_id INT NOT NULL REFERENCES staff_credentials(staff_id) ON DELETE CASCADE,
    CONSTRAINT unique_staff_id UNIQUE (staff_id)
);

-- +goose Down
DROP TABLE staff_info;