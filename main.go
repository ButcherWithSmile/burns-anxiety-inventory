package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/Hatef-PR/Burns-Anxiety-Inventory/PrintLastResults"
	"github.com/Hatef-PR/Burns-Anxiety-Inventory/SymptomsOfPhysicalDiscomfort"
	"github.com/Hatef-PR/Burns-Anxiety-Inventory/WorryingFeelings"
	"github.com/Hatef-PR/Burns-Anxiety-Inventory/WorryingThoughts"


	"database/sql"
	_ "github.com/lib/pq"
)

const (
	host     = // Enter your hostname
	port     = // Enter your port number
	user     = // Enter your PostgreSQL username
	password = // Enter your password
	dbname   = // Enter your database name
)

type Patient struct {
	name   string
	age    int
	date   time.Time
	score  int
	result string
}

func main() {
	var p Patient

	fmt.Println("Please enter your name: ")
	input := bufio.NewReader(os.Stdin)
	name, _ := input.ReadString('\n')
	p.name = strings.TrimSpace(name)

	fmt.Println("Please enter you age: ")
	fmt.Scanln(&p.age)

	p.date = time.Now()
	dateLayout := p.date.Format("2006-01-02 ; 15:04:05")

	p.score = worryingFeelings.WorryingFeelings() + worryingThoughts.WorryingThoughts() +
		SymptomsOfPhysicalDiscomfort.SymptomsOfPhysicalDiscomfort()

	switch {
	case p.score >= 0 && p.score <= 4:
		p.result = "Little anxiety"
	case p.score >= 5 && p.score <= 10:
		p.result = "The border of anxiety"
	case p.score >= 11 && p.score <= 20:
		p.result = "Mild anxiety"
	case p.score >= 21 && p.score <= 30:
		p.result = "Moderate anxiety"
	case p.score >= 31 && p.score <= 50:
		p.result = "Severe anxiety"
	case p.score >= 51 && p.score <= 99:
		p.result = "Extreme anxiety - Panic"
	}

	fmt.Printf("The score for %s, %d years old, at %q is: %d\n", p.name, p.age,
		dateLayout, p.score)
	fmt.Printf("Your anxiety inventory is: %q.\n", p.result)

	dataSource := fmt.Sprintf("host=%s port=%d user=%s password=%d dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", dataSource)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	sqlStatement := `
        INSERT INTO patient (name, age, date, score, result)
        VALUES ($1, $2, $3, $4, $5)
        RETURNING id`
	id := 1
	err = db.QueryRow(sqlStatement, p.name, p.age, p.date, p.score, p.result).Scan(&id)
	if err != nil {
		panic(err)
	}

	PrintLastResults.PrintLastResults(p.name)

	fmt.Println("Press any key to exit...")
	var anyKey string
	fmt.Scanf("%s", &anyKey)
}
