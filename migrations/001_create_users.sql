
CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    full_name TEXT NOT NULL,
    email TEXT NOT NULL,
    age INT NOT NULL
);
