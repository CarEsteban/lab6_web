package internal

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	_ "github.com/lib/pq"
)

// DB es la instancia global de la base de datos.
// @Description Instancia global que almacena la conexión a la base de datos.
var DB *sql.DB

// InitDB inicializa la conexión a la base de datos.
// @Summary Inicializa la base de datos
// @Description Conecta a la base de datos usando las variables de entorno definidas y reintenta la conexión hasta 5 veces.
// @Tags Database
// @Produce plain
// @Success 200 {string} string "Conexión exitosa a la base de datos"
// @Failure 500 {string} string "No se pudo conectar a la base de datos después de varios intentos"
// @Router / [ignore]
func InitDB() error {
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		dbUser, dbPassword, dbHost, dbPort, dbName)

	var err error
	for i := 0; i < 5; i++ {
		DB, err = sql.Open("postgres", connStr)
		if err == nil {
			// Se realiza un ping para verificar que la conexión es exitosa
			if pingErr := DB.Ping(); pingErr == nil {
				fmt.Println("Conexión exitosa a la base de datos")
				return nil
			}
			// Cerrar la conexión si falla el ping
			DB.Close()
		}
		fmt.Printf("Intento %d de conexión fallido: %v\n", i+1, err)
		time.Sleep(3 * time.Second)
	}
	return fmt.Errorf("no se pudo conectar a la base de datos después de varios intentos: %v", err)
}

