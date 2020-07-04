-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE TABLE tasks (
    id int NOT NULL,
    title varchar(255) NOT NULL,
    content text,
    status int NOT NULL,
    PRIMARY KEY(id)
);

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP TABLE tasks;