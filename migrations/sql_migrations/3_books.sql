-- +migrate StatementBegin

-- +migrate Up
CREATE TABLE books (
    id              SERIAL PRIMARY KEY,
    title           VARCHAR(255) NOT NULL,
    description     VARCHAR(500),
    image_url       VARCHAR(500),
    release_year    INTEGER,
    price           INTEGER,
    total_page      INTEGER,
    thickness       VARCHAR(50),
    category_id     INTEGER REFERENCES categories(id) ON DELETE CASCADE,
    created_at      TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    created_by      VARCHAR(100),
    modified_at     TIMESTAMP,
    modified_by     VARCHAR(100)
);

-- +migrate Down
DROP TABLE IF EXISTS books;

-- +migrate StatementEnd
