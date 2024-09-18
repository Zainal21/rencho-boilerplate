-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS users
(
    id INT              NOT NULL,
    uid    VARCHAR(255) NOT NULL,
    firebase_uid        VARCHAR(255)          NULL,
    email               VARCHAR(255)          NOT NULL,
    name                VARCHAR(255)          NOT NULL,
    phone               VARCHAR(255)          NOT NULL,
    PRIMARY KEY (id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS users;
-- +goose StatementEnd
