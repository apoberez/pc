-- CATEGORY TABLE
DROP TABLE IF EXISTS category;
CREATE TABLE category (
  uid   UUID NOT NULL PRIMARY KEY,
  title TEXT NOT NULL,
  slug  TEXT NOT NULL,
  image TEXT NOT NULL
);

-- MANUFACTURER TABLE
DROP TABLE IF EXISTS manufacturer CASCADE;
CREATE TABLE manufacturer (
  uid   UUID NOT NULL PRIMARY KEY,
  title TEXT NOT NULL
);

-- PRODUCT TABLE
DROP TABLE IF EXISTS product;
CREATE TABLE product (
  uid           UUID  NOT NULL PRIMARY KEY,
  title         TEXT  NOT NULL,
  manufacturer  UUID  NOT NULL REFERENCES manufacturer (uid) ON DELETE CASCADE,
  specification JSONB NOT NULL
)