-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE TABLE kanbans (
    id int NOT NULL,
    title varchar(255),
    PRIMARY KEY(id)
);

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP TABLE kanbans;