-- create product table
CREATE TABLE IF NOT EXISTS product (
  id SERIAL PRIMARY KEY,
  name VARCHAR,
  sku VARCHAR,
  category VARCHAR,
  price FLOAT,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);