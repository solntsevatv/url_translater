CREATE TABLE links
(
    id            SERIAL PRIMARY KEY,
    long_url      VARCHAR(255) NOT NULL UNIQUE,
    short_url     VARCHAR(255) NOT NULL UNIQUE
);