package main

import (
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"lab6/internal"
)

// @title La Liga Tracker API
// @version 1.0
// @description API para gestionar los partidos de La Liga.
// @contact.name Esteban Carcamo
// @contact.email car23016@uvg.edu.gt
// @host localhost:8080
// @BasePath /api

// getMatches godoc
// @Summary Obtiene todos los partidos
// @Description Retorna todos los partidos almacenados en la base de datos.
// @Tags Matches
// @Produce json
// @Success 200 {array} internal.Match
// @Failure 500 {object} map[string]string
// @Router /matches [get]
func getMatches(c *gin.Context) {
	matches, err := internal.GetMatches()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, matches)
}

// getMatchID godoc
// @Summary Obtiene un partido por ID
// @Description Retorna el partido cuyo ID se especifica en la ruta.
// @Tags Matches
// @Produce json
// @Param id path int true "ID del partido"
// @Success 200 {object} internal.Match
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /matches/{id} [get]
func getMatchID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	match, err := internal.GetMatchByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "No se encontró el partido"})
		return
	}
	c.JSON(http.StatusOK, match)
}

// createMatch godoc
// @Summary Crea un nuevo partido
// @Description Crea un partido nuevo a partir de los datos enviados en el body.
// @Tags Matches
// @Accept json
// @Produce json
// @Param match body object true "Datos del partido" 
// @Success 201 {object} map[string]int "ID del partido creado"
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /matches [post]
func createMatch(c *gin.Context) {
	var requestBody struct {
		HomeTeam  string `json:"homeTeam"`
		AwayTeam  string `json:"awayTeam"`
		MatchDate string `json:"matchDate"`
	}

	// Se realiza el binding del JSON enviado
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
		return
	}

	// Se parsea la fecha usando el layout "2006-01-02"
	layout := "2006-01-02"
	parsedDate, err := time.Parse(layout, requestBody.MatchDate)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Fecha inválida, use formato YYYY-MM-DD"})
		return
	}

	// Se construye el objeto Match con la fecha parseada
	match := internal.Match{
		HomeTeam:  requestBody.HomeTeam,
		AwayTeam:  requestBody.AwayTeam,
		MatchDate: parsedDate,
	}

	// Se inserta el partido en la base de datos
	newID, err := internal.CreateMatch(match)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Se retorna el ID del partido creado
	c.JSON(http.StatusCreated, gin.H{"id": newID})
}

// updateMatch godoc
// @Summary Actualiza un partido existente
// @Description Actualiza los datos de un partido existente usando el ID de la ruta.
// @Tags Matches
// @Accept json
// @Produce json
// @Param id path int true "ID del partido"
// @Param match body object true "Datos del partido"
// @Success 200 {object} map[string]string "Mensaje de éxito"
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /matches/{id} [put]
func updateMatch(c *gin.Context) {
	// Se obtiene el ID del partido de la ruta
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var requestBody struct {
		HomeTeam  string `json:"homeTeam"`
		AwayTeam  string `json:"awayTeam"`
		MatchDate string `json:"matchDate"`
	}

	// Se realiza el binding del JSON enviado
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
		return
	}

	// Se parsea la fecha con el layout "2006-01-02"
	layout := "2006-01-02"
	parsedDate, err := time.Parse(layout, requestBody.MatchDate)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Fecha inválida, use formato YYYY-MM-DD"})
		return
	}

	// Se construye el objeto Match con el ID y la fecha parseada
	match := internal.Match{
		ID:        id,
		HomeTeam:  requestBody.HomeTeam,
		AwayTeam:  requestBody.AwayTeam,
		MatchDate: parsedDate,
	}

	// Se actualiza el partido en la base de datos
	if err := internal.UpdateMatch(match); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Partido actualizado correctamente"})
}

// deleteMatch godoc
// @Summary Elimina un partido
// @Description Elimina un partido de la base de datos según el ID proporcionado.
// @Tags Matches
// @Produce json
// @Param id path int true "ID del partido"
// @Success 200 {object} map[string]string "Mensaje de éxito"
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /matches/{id} [delete]
func deleteMatch(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	if err := internal.DeleteMatch(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Partido eliminado"})
}

// updateGoals godoc
// @Summary Incrementa los goles del partido
// @Description Incrementa en 1 el campo goals_match para el partido especificado.
// @Tags Matches
// @Produce json
// @Param id path int true "ID del partido"
// @Success 200 {object} map[string]string "Mensaje de éxito"
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /matches/{id}/goals [patch]
func updateGoals(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	if err := internal.UpdateGoals(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Gol incrementado correctamente"})
}

// updateYellowCards godoc
// @Summary Incrementa las tarjetas amarillas del partido
// @Description Incrementa en 1 el campo yellow_cards_match para el partido especificado.
// @Tags Matches
// @Produce json
// @Param id path int true "ID del partido"
// @Success 200 {object} map[string]string "Mensaje de éxito"
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /matches/{id}/yellowcards [patch]
func updateYellowCards(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	if err := internal.UpdateYellowCards(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Tarjeta amarilla incrementada correctamente"})
}

// updateRedCards godoc
// @Summary Incrementa las tarjetas rojas del partido
// @Description Incrementa en 1 el campo red_cards_matchs para el partido especificado.
// @Tags Matches
// @Produce json
// @Param id path int true "ID del partido"
// @Success 200 {object} map[string]string "Mensaje de éxito"
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /matches/{id}/redcards [patch]
func updateRedCards(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	if err := internal.UpdateRedCards(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Tarjeta roja incrementada correctamente"})
}

// updateExtraTime godoc
// @Summary Establece tiempo extra para el partido
// @Description Activa el campo extra_time (lo establece en TRUE) para el partido especificado.
// @Tags Matches
// @Produce json
// @Param id path int true "ID del partido"
// @Success 200 {object} map[string]string "Mensaje de éxito"
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /matches/{id}/extratime [patch]
func updateExtraTime(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	if err := internal.UpdateExtraTime(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Tiempo extra incrementado correctamente"})
}

func main() {
	// Inicializa la conexión a la base de datos
	if err := internal.InitDB(); err != nil {
		log.Fatalf("Error al conectar a la base de datos: %v", err)
	}

	router := gin.Default()

	// Servir el archivo HTML en la raíz
	router.StaticFile("/", "./LaLigaTracker.html")

	api := router.Group("/api")
	{
		api.GET("/matches", getMatches)
		api.GET("/matches/:id", getMatchID)
		api.POST("/matches", createMatch)
		api.PUT("/matches/:id", updateMatch)
		api.DELETE("/matches/:id", deleteMatch)
		api.PATCH("/matches/:id/goals", updateGoals)
		api.PATCH("/matches/:id/yellowcards", updateYellowCards)
		api.PATCH("/matches/:id/redcards", updateRedCards)
		api.PATCH("/matches/:id/extratime", updateExtraTime)
	}

	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Error al iniciar el servidor: %v", err)
	}
}

