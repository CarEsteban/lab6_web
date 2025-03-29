package main


import (
  "log"
  "net/http"
  "strconv"
  "github.com/gin-gonic/gin"
  "lab6/internal"
  "time"
)

//FUNCIONES para todas las llamadas
func getMatches(c *gin.Context) {
  
  matches, err := internal.GetMatches()
  if err != nil {
    c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
    return
  }
  c.JSON(http.StatusOK, matches)
}

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

func createMatch(c *gin.Context) {
    var requestBody struct {
        HomeTeam  string `json:"homeTeam"`
        AwayTeam  string `json:"awayTeam"`
        MatchDate string `json:"matchDate"` 
      }

    // 2. Bindeamos el JSON que manda el frontend
    if err := c.ShouldBindJSON(&requestBody); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
        return
    }

    // 3. Parseamos la fecha usando el layout "2006-01-02"
    layout := "2006-01-02"
    parsedDate, err := time.Parse(layout, requestBody.MatchDate)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Fecha inválida, use formato YYYY-MM-DD"})
        return
    }

    // 4. Construimos el objeto Match con la fecha parseada
    match := internal.Match{
        HomeTeam:  requestBody.HomeTeam,
        AwayTeam:  requestBody.AwayTeam,
        MatchDate: parsedDate, // se guarda como time.Time
    }

    // 5. Llamamos a la función que inserta en la DB
    newID, err := internal.CreateMatch(match)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    // 6. Respondemos con el ID del partido recién creado
    c.JSON(http.StatusCreated, gin.H{"id": newID})
}


func updateMatch(c *gin.Context) {
    // 1. Obtener el ID desde la URL
    idParam := c.Param("id")
    id, err := strconv.Atoi(idParam)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
        return
    }

    // 2. Definir un struct que reciba la fecha como string
    var requestBody struct {
        HomeTeam  string `json:"homeTeam"`
        AwayTeam  string `json:"awayTeam"`
        MatchDate string `json:"matchDate"` // Recibe "2025-04-01"
    }

    // 3. Bindeamos el JSON que manda el frontend
    if err := c.ShouldBindJSON(&requestBody); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
        return
    }

    // 4. Parseamos la fecha usando "2006-01-02" (YYYY-MM-DD)
    layout := "2006-01-02"
    parsedDate, err := time.Parse(layout, requestBody.MatchDate)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Fecha inválida, use formato YYYY-MM-DD"})
        return
    }

    // 5. Construimos el objeto Match con el ID y la fecha parseada
    match := internal.Match{
        ID:        id,
        HomeTeam:  requestBody.HomeTeam,
        AwayTeam:  requestBody.AwayTeam,
        MatchDate: parsedDate,
    }

    // 6. Llamamos a la función que actualiza en la DB
    if err := internal.UpdateMatch(match); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    // 7. Respondemos con un mensaje de éxito
    c.JSON(http.StatusOK, gin.H{"message": "Partido actualizado correctamente"})
}


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

func updateGoals(c *gin.Context){


	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	// Llamamos a la función que incrementa el gol en la base de datos.
	if err := internal.UpdateGoals(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Gol incrementado correctamente"})


}


func updateYellowCards(c *gin.Context){


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

	c.JSON(http.StatusOK, gin.H{"message": "Tarjeta amarrilla incrementada correctamente"})


}

func updateRedCards(c *gin.Context){


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

func updateExtraTime(c *gin.Context){


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

	c.JSON(http.StatusOK, gin.H{"message": "Tiempo extra incrementada correctamente"})


}

func main() {
  //cone con la BD
  if err := internal.InitDB(); err != nil {
    log.Fatalf("Error al conectar a la base de datos: %v", err)
  }

  router := gin.Default()
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
