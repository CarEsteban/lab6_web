package internal

import (	
	"fmt"
	"time"
	// Asegúrate de importar el driver de PostgreSQL
	_ "github.com/lib/pq"
)


// Match representa la estructura de un partido de La Liga.
// @Description Objeto que modela un partido, incluyendo equipos y fecha.
type Match struct {
	ID        int       `json:"id"`
	HomeTeam  string    `json:"homeTeam"`
	AwayTeam  string    `json:"awayTeam"`
	MatchDate time.Time `json:"matchDate"`
}

// GetMatches obtiene todos los partidos de la base de datos.
// @Summary Obtiene todos los partidos
// @Description Realiza una consulta a la tabla "matches" y retorna una lista de partidos.
// @Success 200 {array} Match "Lista de partidos"
// @Failure 500 {object} map[string]string "Error interno"
// @Router /matches [get]
func GetMatches() ([]Match, error) {
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

// GetMatchByID obtiene un partido según su ID.
// @Summary Obtiene un partido por ID
// @Description Realiza una consulta para obtener un partido específico mediante su ID.
// @Param id path int true "ID del partido"
// @Success 200 {object} Match "Partido encontrado"
// @Failure 404 {object} map[string]string "Partido no encontrado"
// @Router /matches/{id} [get]
func GetMatchByID(id int) (Match, error) {
	var m Match
	if err := DB.QueryRow("SELECT id, home_team, away_team, match_date FROM matches WHERE id = $1", id).
		Scan(&m.ID, &m.HomeTeam, &m.AwayTeam, &m.MatchDate); err != nil {
		return m, err
	}
	return m, nil
}

// CreateMatch inserta un nuevo partido en la base de datos.
// @Summary Crea un nuevo partido
// @Description Inserta en la tabla "matches" un nuevo registro y retorna su ID.
// @Param m body Match true "Objeto Match sin ID"
// @Success 201 {int} int "ID del partido creado"
// @Failure 500 {object} map[string]string "Error al crear el partido"
// @Router /matches [post]
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

// UpdateMatch actualiza un partido existente en la base de datos.
// @Summary Actualiza un partido
// @Description Actualiza los campos home_team, away_team y match_date para un partido dado.
// @Param m body Match true "Objeto Match con ID y datos actualizados"
// @Success 200 {object} map[string]string "Partido actualizado correctamente"
// @Failure 500 {object} map[string]string "Error al actualizar el partido"
// @Router /matches/{id} [put]
func UpdateMatch(m Match) error {
	query := `
        UPDATE matches
        SET home_team = $1, away_team = $2, match_date = $3
        WHERE id = $4
    `
	_, err := DB.Exec(query, m.HomeTeam, m.AwayTeam, m.MatchDate, m.ID)
	return err
}

// DeleteMatch elimina un partido de la base de datos.
// @Summary Elimina un partido
// @Description Elimina el registro de la tabla "matches" correspondiente al ID proporcionado.
// @Param id path int true "ID del partido"
// @Success 200 {object} map[string]string "Partido eliminado"
// @Failure 500 {object} map[string]string "Error al eliminar el partido"
// @Router /matches/{id} [delete]
func DeleteMatch(id int) error {
	query := `
        DELETE FROM matches
        WHERE id = $1
    `
	_, err := DB.Exec(query, id)
	return err
}

// UpdateGoals incrementa en 1 el campo goals_match para el partido dado.
// @Summary Incrementa goles del partido
// @Description Incrementa el valor de "goals_match" en 1 para el partido especificado.
// @Param id path int true "ID del partido"
// @Success 200 {object} map[string]string "Gol incrementado correctamente"
// @Failure 500 {object} map[string]string "Error al incrementar goles"
// @Router /matches/{id}/goals [patch]
func UpdateGoals(id int) error {
	query := "UPDATE matches SET goals_match = goals_match + 1 WHERE id = $1"
	_, err := DB.Exec(query, id)
	if err != nil {
		return fmt.Errorf("error al incrementar goles: %v", err)
	}
	return nil
}

// UpdateYellowCards incrementa en 1 el campo yellow_cards_match para el partido dado.
// @Summary Incrementa tarjetas amarillas
// @Description Incrementa el valor de "yellow_cards_match" en 1 para el partido especificado.
// @Param id path int true "ID del partido"
// @Success 200 {object} map[string]string "Tarjeta amarilla incrementada correctamente"
// @Failure 500 {object} map[string]string "Error al incrementar tarjeta amarilla"
// @Router /matches/{id}/yellowcards [patch]
func UpdateYellowCards(id int) error {
	query := "UPDATE matches SET yellow_cards_match = yellow_cards_match + 1 WHERE id = $1"
	_, err := DB.Exec(query, id)
	if err != nil {
		return fmt.Errorf("error al incrementar tarjeta amarilla: %v", err)
	}
	return nil
}

// UpdateRedCards incrementa en 1 el campo red_cards_match para el partido dado.
// @Summary Incrementa tarjetas rojas
// @Description Incrementa el valor de "red_cards_match" en 1 para el partido especificado.
// @Param id path int true "ID del partido"
// @Success 200 {object} map[string]string "Tarjeta roja incrementada correctamente"
// @Failure 500 {object} map[string]string "Error al incrementar tarjeta roja"
// @Router /matches/{id}/redcards [patch]
func UpdateRedCards(id int) error {
	query := "UPDATE matches SET red_cards_match = red_cards_match + 1 WHERE id = $1"
	_, err := DB.Exec(query, id)
	if err != nil {
		return fmt.Errorf("error al incrementar tarjeta roja: %v", err)
	}
	return nil
}

// UpdateExtraTime establece en TRUE el campo extra_time para el partido dado.
// @Summary Activa tiempo extra
// @Description Establece el valor de "extra_time" en TRUE para el partido especificado.
// @Param id path int true "ID del partido"
// @Success 200 {object} map[string]string "Tiempo extra establecido correctamente"
// @Failure 500 {object} map[string]string "Error al establecer tiempo extra"
// @Router /matches/{id}/extratime [patch]
func UpdateExtraTime(id int) error {
	query := "UPDATE matches SET extra_time = TRUE WHERE id = $1"
	_, err := DB.Exec(query, id)
	if err != nil {
		return fmt.Errorf("error al establecer tiempo extra: %v", err)
	}
	return nil
}

