CREATE TABLE test (
  id BIGSERIAL NOT NULL PRIMARY KEY,
  name VARCHAR(256),
  created_at TIMESTAMP default NOW(),
  updated_at TIMESTAMP DEFAULT now()
);