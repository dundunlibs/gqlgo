CREATE TABLE users (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  name VARCHAR(255)
);
CREATE TABLE posts (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  title VARCHAR(255),
  body TEXT,
  author_id INTEGER NOT NULL,
  FOREIGN KEY (author_id) REFERENCES users(id)
);

INSERT INTO
  users (name)
VALUES
  ('John Doe'),
  ('Scarlet Smith');

INSERT INTO
  posts (title, body, author_id)
VALUES
  ('Hello World!', 'Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua', 1),
  ('Section 1.10.32 of "de Finibus Bonorum et Malorum"', 'Sed ut perspiciatis unde omnis iste natus error sit voluptatem accusantium doloremque laudantium, totam rem aperiam, eaque ipsa quae ab illo inventore veritatis et quasi architecto beatae vitae dicta sunt explicabo', 2);
