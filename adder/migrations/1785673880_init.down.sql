CREATE TABLE IF NOT EXISTS radio_stations (
    id UUID PRIMARY KEY DEFAULT uuidv7(),
    name VARCHAR(100) NOT NULL,
    url VARCHAR(500) NOT NULL UNIQUE,
    genre VARCHAR(50) NOT NULL,
    country VARCHAR(50),
    language VARCHAR(50),
    is_active BOOLEAN DEFAULT true,
    popular_rating DECIMAL(3,2) DEFAULT 0.00,
    total_plays INTEGER DEFAULT 0,
    added_by VARCHAR(100),
    added_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);


CREATE TABLE IF NOT EXISTS station_plays (
    id UUID PRIMARY KEY DEFAULT uuidv7(),
    station_id INTEGER REFERENCES radio_stations(id) ON DELETE CASCADE,
    session_id VARCHAR(100) NOT NULL,
    played_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    rating INTEGER CHECK (rating >= 1 AND rating <= 5),
    UNIQUE(station_id, session_id)
);