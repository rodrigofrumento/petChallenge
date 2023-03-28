CREATE TABLE pets (
    id SERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    age INTEGER,
    description TEXT,
    city VARCHAR(50) NOT NULL
);