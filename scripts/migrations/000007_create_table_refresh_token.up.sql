CREATE TABLE IF NOT EXISTS refresh_tokens (
    id INT AUTO_INCREMENT PRIMARY KEY,
    user_id BIGINT NOT NULL, 
    refresh_token TEXT NOT NULL,
    expired_at TIMESTAMP NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    created_by LONGTEXT NOT NULL,
    updated_by LONGTEXT NOT NULL,
    CONSTRAINT fk_user_id_refresh_tokens FOREIGN KEY (user_id) REFERENCES users(id)
);