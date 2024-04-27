package connection

import (
    "database/sql"
    "log"
    _ "github.com/lib/pq"
)

func ConnectDB() (*sql.DB, error) {
	db, err := sql.Open("postgres", "user=postgres password=postgres dbname=ktaxes sslmode=disable") 

    if err != nil {
        return nil, err
    }
    log.Println("Database connection successful")
    return db, nil
}
