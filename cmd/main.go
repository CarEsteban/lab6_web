package main


import (
  "log"
  "net/http"
  "strconv"
  "github.com/gin-gonic/gin"
  "lab6/internal"
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
  var newMatch internal.Match
  if err := c.ShouldBindJSON(&newMatch); err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
    return
  }

  newID, err := internal.CreateMatch(newMatch)
  if err != nil {
    c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
    return
  }

  c.JSON(http.StatusCreated, gin.H{"id": newID})
}

func updateMatch(c *gin.Context) {
  idParam := c.Param("id")
  id, err := strconv.Atoi(idParam)
  if err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
    return
  }

  var updatedMatch internal.Match
  if err := c.ShouldBindJSON(&updatedMatch); err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
    return
  }

  updatedMatch.ID = id
  if err := internal.UpdateMatch(updatedMatch); err != nil {
    c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
    return
  }
  c.JSON(http.StatusOK, gin.H{"message": "Partido actualizado"})
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
  }
  
  if err := router.Run(":8080"); err != nil {
    log.Fatalf("Error al iniciar el servidor: %v", err)
  }

}
