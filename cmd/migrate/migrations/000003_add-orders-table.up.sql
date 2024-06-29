DROP TYPE IF EXISTS status;

CREATE TYPE status AS ENUM ('pending', 'completed', 'cancelled');

CREATE TABLE IF NOT EXISTS orders (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL,
    total DECIMAL(10, 2) NOT NULL,
    type status NOT NULL DEFAULT 'pending',
    address TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_users
        FOREIGN KEY(user_id)
            REFERENCES users(id)
);