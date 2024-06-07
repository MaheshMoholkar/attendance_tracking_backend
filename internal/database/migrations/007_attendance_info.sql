-- +goose Up
CREATE TABLE attendance_info (
    attendance_id SERIAL PRIMARY KEY,
    attendance_table_name TEXT NOT NULL,
    attendance_month_year TEXT NOT NULL,
    class_id INT NOT NULL REFERENCES class_info(class_id),
    division_id INT NOT NULL REFERENCES division_info(division_id)
);

-- +goose Down
DROP TABLE attendance_info;
