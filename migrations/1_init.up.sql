CREATE TABLE if not exists offices
(
    uuid            uuid         not null,
    office_name     varchar(50)  not null,
    address         varchar(100) not null,
    created_at      timestamp default now(),
    updated_at      timestamp default now(),

    PRIMARY KEY (uuid)
);
CREATE INDEX IF NOT EXISTS idx_offices_name ON offices (name);

CREATE TABLE IF NOT EXISTS users
(
    uuid        uuid                default uuid_generate_v4(),
    firstname   varchar(50)         NOT NULL,
    lastname    varchar(50),
    office_uuid uuid                NOT NULL,
    office_name varchar(50)         NOT NULL,
    email       varchar(255) unique NOT NULL,
    password    varchar(255)        NOT NULL,
    created_at  timestamp default now(),
    updated_at  timestamp default now(),

    PRIMARY KEY (uuid),
    FOREIGN KEY (office_uuid) REFERENCES offices (uuid)
);
CREATE INDEX IF NOT EXISTS idx_email ON users (email);