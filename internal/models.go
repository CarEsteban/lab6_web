package internal

import (
  "time"
)


type Match struct {
  ID        int       `json:"id"`
	HomeTeam  string    `json:"homeTeam"`
	AwayTeam  string    `json:"awayTeam"`
	MatchDate time.Time `json:"matchDate"`  
}

//funcionar para obtener todos los partidos
func GetMatches() ([]Match,error){
  rows, err := DB.Query("SELECT id, home_team, away_team, match_date FROM matches")
  if err != nil {
    return nil, err
  }

  defer rows.Close()

  var matches []Match
  for rows.Next() {
    var m Match
    if err := rows.Scan(&m.ID, &m.HomeTeam, &m.AwayTeam, &m.MatchDate); err != nil {
      return nil, err
    }
    matches = append(matches, m)
  }
  return matches, nil
}

//fncion para obtener un partido mediante su id
func GetMatchByID(id int) (Match, error) {
  var m Match
  if err := DB.QueryRow("SELECT id, home_team, away_team, match_date FROM matches WHERE id = $1", id).Scan(&m.ID, &m.HomeTeam, &m.AwayTeam, &m.MatchDate); err != nil {
    return m, err
  }
  return m, nil
}

//funcion para crear un partido
func CreateMatch(m Match) (int, error) {
    query := `
        INSERT INTO matches (home_team, away_team, match_date)
        VALUES ($1, $2, $3)
        RETURNING id
    `
    var newID int
    err := DB.QueryRow(query, m.HomeTeam, m.AwayTeam, m.MatchDate).Scan(&newID)
    return newID, err
}

//funcion para actualizar un partido existente
func UpdateMatch(m Match) error {
    query := `
        UPDATE matches
        SET home_team=$1, away_team=$2, match_date=$3
        WHERE id=$4
    `
    _, err := DB.Exec(query, m.HomeTeam, m.AwayTeam, m.MatchDate, m.ID)
    return err
}

//funcion para eliminar un partido
func DeleteMatch(id int) error {
    query := `
        DELETE FROM matches
        WHERE id=$1
    `
    _, err := DB.Exec(query, id)
    return err
}
