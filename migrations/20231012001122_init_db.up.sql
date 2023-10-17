START TRANSACTION;

CREATE TABLE users
(
    id         BIGINT AUTO_INCREMENT PRIMARY KEY,
    name       VARCHAR(64)                         NOT NULL,
    status     VARCHAR(16)                         NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP NOT NULL
);

CREATE TABLE user_accounts
(
    id         BIGINT AUTO_INCREMENT PRIMARY KEY,
    user_id    BIGINT                              NOT NULL,
    name       VARCHAR(64)                         NOT NULL,
    status     VARCHAR(16)                         NOT NULL,
    balance    DOUBLE PRECISION                    NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP NOT NULL
);

ALTER TABLE user_accounts
    ADD CONSTRAINT user_accounts_user_id_fkey FOREIGN KEY (user_id) REFERENCES users (id);

COMMIT;