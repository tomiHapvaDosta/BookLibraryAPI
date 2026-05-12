-- +goose Up
CREATE TABLE 
    books(
        id UUID PRIMARY KEY,
        title TEXT UNIQUE,
        author TEXT NOT NULL,
        genre TEXT NOT NULL,
        publication_year INT NOT NULL,
        created_at TIMESTAMP NOT NULL,
        updated_at TIMESTAMP NOT NULL
    );

-- +goose Down
DROP TABLE books;