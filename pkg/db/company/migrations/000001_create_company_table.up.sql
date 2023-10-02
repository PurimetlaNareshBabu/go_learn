CREATE TABLE companies (
    id INT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
    created_by INT UNSIGNED,
    created_at BIGINT,
    updated_at BIGINT,
    is_deleted BOOLEAN,
    FOREIGN KEY (created_by) REFERENCES users(id)
);

CREATE TABLE roles (
    id INT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
    created_by INT UNSIGNED,
    created_at BIGINT,
    updated_at BIGINT,
    is_deleted BOOLEAN,
    company_id INT UNSIGNED NOT NULL,
    FOREIGN KEY (created_by) REFERENCES users(id),
    FOREIGN KEY (company_id) REFERENCES companies(id)
);