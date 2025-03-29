-- Crear la tabla
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

-- Insertar datos iniciales
INSERT INTO matches (home_team, away_team, match_date)
VALUES
  ('Barcelona', 'Real Madrid', '2026-04-01'),
  ('Atletico Madrid', 'Sevilla', '2025-04-02'),
  ('Valencia', 'Villarreal', '2025-04-03'),
  ('Real Sociedad', 'Athletic Club', '2025-04-04'),
  ('Betis', 'Getafe', '2025-04-05'),
  ('Espanyol', 'Celta de Vigo', '2025-04-06'),
  ('Levante', 'Real Valladolid', '2025-04-07'),
  ('Granada', 'Mallorca', '2025-04-08'),
  ('Cadiz', 'Elche', '2025-04-09'),
  ('Almeria', 'Osasuna', '2025-04-10');

