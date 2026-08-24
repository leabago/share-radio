CREATE TABLE radio_stations (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    url VARCHAR(500) NOT NULL UNIQUE, -- prevent duplicate URLs
    genre VARCHAR(50) NOT NULL,
    country VARCHAR(50),
    language VARCHAR(50),
    is_active BOOLEAN DEFAULT true,
    popular_rating DECIMAL(3,2) DEFAULT 0.00, -- rating from 0-5
    total_plays INTEGER DEFAULT 0,
    added_by VARCHAR(100), -- optional: track who added it (IP or nickname)
    added_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);