package db

import (
	"database/sql"
	"fmt"
)

var db *sql.DB

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "********"
	dbname   = "slackbotdb"
)

func DBConnecting(env string) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, _ = sql.Open("postgres", psqlInfo)

	defer db.Close()

	err := db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")

}

//Example
func sampleDataEntry(db *sql.DB) {
	sqlStatement := `
	INSERT INTO users (age, email, first_name, last_name)
	VALUES ($1, $2, $3, $4)
	RETURNING id`
	id := 0
	db.QueryRow(sqlStatement, 30, "abc@g.co", "bac", "abc").Scan(&id)
	fmt.Println("New Record ID is:", id)
}

func GetClient() *sql.DB {
	return db
}
