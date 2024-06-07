-- +goose Up
CREATE TABLE class_info (
    class_id SERIAL PRIMARY KEY,
    className TEXT NOT NULL
);

-- +goose Down
DROP TABLE class_info;