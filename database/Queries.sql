create database shopping_service_db;

CREATE TABLE users
(
    id         serial PRIMARY KEY,
    first_name VARCHAR(55),
    last_name  VARCHAR(55),
    login      VARCHAR(55),
    password   VARCHAR(55),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMP
);

CREATE TABLE items
(
    id         serial PRIMARY KEY,
    title VARCHAR(55),
    price  DOUBLE PRECISION,
    country     VARCHAR(55),
    description   VARCHAR(55),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMP
);

select *
from items;

select *
from users;
