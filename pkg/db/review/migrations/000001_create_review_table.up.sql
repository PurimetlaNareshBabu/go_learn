CREATE TABLE reviews (
    id INT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
    content TEXT NOT NULL,
    created_at BIGINT,
    updated_at BIGINT,
    created_by_id INT UNSIGNED NOT NULL,
    company_id INT UNSIGNED NOT NULL,
    role_id INT UNSIGNED NOT NULL,
    title VARCHAR(255),
    FOREIGN KEY (created_by_id) REFERENCES users(id),
    FOREIGN KEY (company_id) REFERENCES companies(id),
    FOREIGN KEY (role_id) REFERENCES roles(id)
);