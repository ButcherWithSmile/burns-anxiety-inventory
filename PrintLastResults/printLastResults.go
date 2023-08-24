package PrintLastResults

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"time"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = 5723245
	dbname   = "BAIT"
)

type Patient struct {
	name   string
	date   time.Time
	age    int
	score  int
	result string
}

func PrintLastResults(name string) {
	dataSource := fmt.Sprintf("host=%s port=%d user=%s password=%d dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", dataSource)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

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

	dateLayout := p.date.Format("2006-01-02 ; 15:04:05")

	fmt.Printf("The last score was %d and result was %q, at %q.\n", p.score, p.result, dateLayout)
}
