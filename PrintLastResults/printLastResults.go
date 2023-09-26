// This package imports the necessary libraries for connecting to a PostgreSQL database, querying data, and formatting dates

package PrintLastResults

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"time"
)

// These constants store the database connection parameters
const (
	host     = // Enter your hostname
	port     = // Enter your port number
	user     = // Enter your PostgreSQL username
	password = // Enter your password
	dbname   = // Enter your database name
)

// This struct represents a patient and their test results
type Patient struct {
	name   string
	date   time.Time
	age    int
	score  int
	result string
}

// This function connects to the PostgreSQL database using the specified connection parameters
func PrintLastResults(name string) {
	dataSource := fmt.Sprintf("host=%s port=%d user=%s password=%d dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", dataSource)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// This code queries the database for the last test results for the given patient name
	// If no records are found, the function prints a message to the console and returns
	// Otherwise, the function scans the results into a Patient struct
	var p Patient
	row := db.QueryRow(`SELECT name, age, date, score, result FROM patient WHERE name=$1 
                                                   ORDER BY date DESC LIMIT 1`, name)
	err = row.Scan(&p.name, &p.age, &p.date, &p.score, &p.result)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			fmt.Println("No records found for", name)
			return
		}
		panic(err)
	}

	// This code formats the date and prints the patient's last test results to the console
	dateLayout := p.date.Format("2006-01-02 ; 15:04:05")

	fmt.Printf("The last score was %d and result was %q, at %q.\n", p.score, p.result, dateLayout)
}
