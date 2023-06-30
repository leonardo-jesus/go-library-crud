CREATE TABLE books (
  id SERIAL PRIMARY KEY,
  name VARCHAR(255) NOT NULL,
  edition INTEGER,
  publication_year INTEGER
);