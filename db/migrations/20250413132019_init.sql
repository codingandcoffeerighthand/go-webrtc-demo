-- +goose Up
-- +goose StatementBegin
CREATE TABLE stream_room (
    wallet_id TEXT NOT NULL,
    stream_name TEXT UNIQUE NOT NULL,
    PRIMARY KEY (wallet_id, stream_name)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE stream_room;
-- +goose StatementEnd