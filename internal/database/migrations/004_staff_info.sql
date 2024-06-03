-- +goose Up
CREATE TABLE staff_info (
    id SERIAL PRIMARY KEY,
    firstName TEXT NOT NULL,
    lastName TEXT NOT NULL,
    email TEXT NOT NULL,
    staff_id INT NOT NULL REFERENCES staff_credentials(staff_id) ON DELETE CASCADE
);

-- +goose Down
DROP TABLE staff;