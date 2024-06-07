-- +goose Up
CREATE TABLE division_info (
    division_id SERIAL PRIMARY KEY,
    divisionName TEXT NOT NULL,
    class_id INT NOT NULL REFERENCES class_info(class_id) ON DELETE CASCADE
);

-- +goose Down
DROP TABLE division_info;