-- +goose Up
-- +goose StatementBegin
CREATE TABLE user (
    id INT PRIMARY KEY AUTO_INCREMENT,
    full_name VARCHAR(255),
    user_name VARCHAR(255),
    email_id VARCHAR(255),
    password VARCHAR(255),
    created_at BIGINT NOT NULL,
    updated_at BIGINT NOT NULL,
    last_login BIGINT
    is_deleted BOOLEAN
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS `user`;
-- +goose StatementEnd
