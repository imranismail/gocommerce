CREATE TABLE users (
  id SERIAL PRIMARY KEY,
  email VARCHAR,
  hashed_password VARCHAR
);