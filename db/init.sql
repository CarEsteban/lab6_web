
/*
========================================================================
SCRIPT DE CREACIÓN E INSERT DE DATOS PARA LA TABLA "matches"
========================================================================

Descripción:
Este script crea la tabla "matches" en PostgreSQL, la cual almacena la información 
de los partidos de La Liga. Además, inserta datos iniciales para poblar la tabla y
facilitar las pruebas en el desarrollo del backend.

Estructura de la Tabla:
  - id                : Identificador único del partido (SERIAL, PRIMARY KEY)
  - home_team         : Nombre del equipo local (VARCHAR(100), NOT NULL)
  - away_team         : Nombre del equipo visitante (VARCHAR(100), NOT NULL)
  - match_date        : Fecha del partido (DATE, NOT NULL)
  - goals_match       : Total de goles anotados en el partido (INT, DEFAULT 0)
  - yellow_cards_match: Total de tarjetas amarillas (INT, DEFAULT 0)
  - red_cards_match   : Total de tarjetas rojas (INT, DEFAULT 0)
  - extra_time        : Indica si se jugó tiempo extra (BOOLEAN, DEFAULT FALSE)

========================================================================
*/

/* Crear la tabla "matches" si no existe */
CREATE TABLE IF NOT EXISTS matches (
    id SERIAL PRIMARY KEY,
    home_team VARCHAR(100) NOT NULL,
    away_team VARCHAR(100) NOT NULL,
    match_date DATE NOT NULL,
    goals_match INT DEFAULT 0,
    yellow_cards_match INT DEFAULT 0,
    red_cards_match INT DEFAULT 0,
    extra_time BOOLEAN DEFAULT FALSE
);

/*========================================================================
   Insertar datos iniciales en la tabla "matches"
========================================================================*/

/*
Se insertan 10 registros de ejemplo con partidos de La Liga.
La fecha debe cumplir con el formato 'YYYY-MM-DD'.
Los campos de goles, tarjetas y tiempo extra utilizarán los valores por defecto.
*/
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

