start transaction;

create table users
(
    id         bigint AUTO_INCREMENT PRIMARY KEY,
    name       varchar(64)                         NOT NULL,
    status     varchar(16)                         NOT NULL,
    created_at timestamp DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at timestamp DEFAULT CURRENT_TIMESTAMP NOT NULL ON UPDATE CURRENT_TIMESTAMP
);

create table user_accounts
(
    id         bigint AUTO_INCREMENT PRIMARY KEY,
    user_id    bigint                              NOT NULL,
    name       varchar(64)                         NOT NULL,
    status     varchar(16)                         NOT NULL,
    balance    double precision                    NOT NULL,
    created_at timestamp DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at timestamp DEFAULT CURRENT_TIMESTAMP NOT NULL ON UPDATE CURRENT_TIMESTAMP
);

alter table user_accounts
    add constraint user_accounts_user_id_fkey foreign key (user_id) REFERENCES users (id);

commit;