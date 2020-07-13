-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE TABLE tasks (
    id int NOT NULL AUTO_INCREMENT,
    kanban_id int NOT NULL,
    title varchar(255) NOT NULL,
    content text,
    status int NOT NULL DEFAULT 1 COMMENT '1: todo, 2: doing, 3: done',
    created_at datetime,
    updated_at datetime,
    deleted_at datetime,
    PRIMARY KEY(id),
    FOREIGN KEY (kanban_id) REFERENCES kanbans(id) ON DELETE CASCADE
);

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP TABLE tasks;