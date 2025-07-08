-- +goose Up
-- +goose StatementBegin
CREATE TABLE todo (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description TEXT 
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE todo;
-- +goose StatementEnd
