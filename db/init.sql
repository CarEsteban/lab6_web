CREATE TABLE IF NOT EXISTS matches (
    id SERIAL PRIMARY KEY,
    home_team VARCHAR(100) NOT NULL,
    away_team VARCHAR(100) NOT NULL,
    match_date DATE NOT NULL,
    home_goals INT DEFAULT 0,
    away_goals INT DEFAULT 0,
    home_yellow_cards INT DEFAULT 0,
    away_yellow_cards INT DEFAULT 0,
    home_red_cards INT DEFAULT 0,
    away_red_cards INT DEFAULT 0,
    extra_time BOOLEAN DEFAULT FALSE
);
