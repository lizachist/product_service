CREATE TABLE products
(
    id          SERIAL PRIMARY KEY,
    name        VARCHAR(50) UNIQUE NOT NULL,
    price       DECIMAL(10, 2)     NOT NULL,
    quantity    INTEGER            NOT NULL,
    description TEXT,
    category_id INTEGER REFERENCES category (id),
    is_active   BOOLEAN            NOT NULL DEFAULT TRUE,
    created_at  TIMESTAMP WITH TIME ZONE    DEFAULT CURRENT_TIMESTAMP,
    updated_at  TIMESTAMP WITH TIME ZONE    DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE category
(
    id         SERIAL PRIMARY KEY,
    name       VARCHAR(50) UNIQUE NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);