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

func PrepareDB() {
	db, err := ConnectDB()
	if err != nil {
		log.Fatal("Error connecting to the database:", err)
	}

	rows, err := db.Query("SELECT EXISTS (SELECT FROM information_schema.tables WHERE table_name = 'tax')")
	if err != nil {
		log.Fatal("Error checking if 'tax' table exists:", err)
	}
	defer rows.Close()

	var exists bool
	for rows.Next() {
		if err := rows.Scan(&exists); err != nil {
			log.Fatal("Error scanning rows:", err)
		}
	}

	if !exists {
		if _, err := db.Exec(`CREATE TABLE tax (
			id SERIAL PRIMARY KEY,
			k_reciept_deduct INTEGER,
			personal_deduct INTEGER,
			active BOOLEAN,
			create_time TIMESTAMP
		)`); err != nil {
			log.Fatal("Error creating 'tax' table:", err)
		}
		log.Println("Table 'tax' created successfully")

		if _, err := db.Exec(`INSERT INTO tax (k_reciept_deduct, personal_deduct, active, create_time)
			VALUES (50000, 60000, true, NOW())`); err != nil {
			log.Fatal("Error inserting initial data into 'tax' table:", err)
		}
		log.Println("Initial data inserted into 'tax' table successfully")
	} else {
		log.Println("Table 'tax' already exists")
	}

	defer db.Close()
}
