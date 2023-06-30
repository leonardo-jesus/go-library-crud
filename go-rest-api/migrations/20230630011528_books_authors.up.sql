CREATE TABLE book_authors (
  book_id INTEGER,
  author_id INTEGER,
  FOREIGN KEY (book_id) REFERENCES books (id),
  FOREIGN KEY (author_id) REFERENCES authors (id),
  PRIMARY KEY (book_id, author_id)
);