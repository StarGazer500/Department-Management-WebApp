package database

import (
	"fmt"
	_ "github.com/lib/pq"
    "database/sql"
	"os"
)

func ConnectTODb() (error) {

	dbPort := os.Getenv("DBPORT")
	dbSslMode := os.Getenv("dbSslMode")
	dbUser := os.Getenv("DBUSER")
	dbPassword := os.Getenv("DBPASSWORD")
	dbName := os.Getenv("DBNAME")
	dbHost := os.Getenv("HOSTDEV")
	// Construct the connection string using the provided parameters
	connStr := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=%s",
		dbUser, dbPassword, dbName, dbHost, dbPort, dbSslMode)

	// Open the connection to the database
	fmt.Println("port",dbPort)
	fmt.Println("connection string", connStr)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return fmt.Errorf("error gconnecting to the database: %v", err)
	}

	// Check if the connection is valid by pinging the database
	if err := db.Ping(); err != nil {
		 return fmt.Errorf("unable to reach the database: %v", err)
	}

	// Assign the database instance to the global PG variable
	// PG := &DbInstance{Db: db}
	fmt.Println("Successfully connected to the database!")

	// Return nil if everything went well
	SetDb(db)
	return nil
}