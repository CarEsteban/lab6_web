package internal

import (
  "database/sql"
  "fmt"
  "time"
  "os"
  _ "github.com/lib/pq"
)

var DB *sql.DB


func InitDB() error {
    dbHost := os.Getenv("DB_HOST")
    dbPort := os.Getenv("DB_PORT")
    dbUser := os.Getenv("DB_USER")
    dbPassword := os.Getenv("DB_PASSWORD")
    dbName := os.Getenv("DB_NAME")

    connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",dbUser, dbPassword, dbHost, dbPort, dbName)

    var err error
    for i := 0; i < 5; i++ {
        DB, err = sql.Open("postgres", connStr)
        if err == nil {
            // Probamos el ping para ver si realmente está conectada
            if pingErr := DB.Ping(); pingErr == nil {
                fmt.Println("Conexión exitosa a la base de datos")
                return nil
            }
            // Cerrar si falla el ping
            DB.Close()
        }
        fmt.Printf("Intento %d de conexión fallido: %v\n", i+1, err)
        time.Sleep(3 * time.Second) // Espera 3 segundos antes de reintentar
    }
    return fmt.Errorf("no se pudo conectar a la base de datos después de varios intentos: %v", err)
}

