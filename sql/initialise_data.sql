CREATE TABLE IF NOT EXISTS tracks (
    id SERIAL PRIMARY KEY,
    source VARCHAR(255) NOT NULL,
    target VARCHAR(255) NOT NULL,
    is_deleted BOOLEAN DEFAULT FALSE
);

CREATE TABLE IF NOT EXISTS signals (
    id SERIAL PRIMARY KEY,
    signal_name VARCHAR(255) NOT NULL,
    elr VARCHAR(50),
    mileage FLOAT8,
    track_id INT,
    is_deleted BOOLEAN DEFAULT FALSE
);