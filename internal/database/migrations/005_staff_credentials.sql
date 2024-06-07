-- +goose Up
CREATE TABLE staff_credentials(
    staff_id INT PRIMARY KEY,
    password_hash TEXT NOT NULL
);

-- +goose Down
DROP TABLE staff_credentials;