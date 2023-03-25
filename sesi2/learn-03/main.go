package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "ssasongko"
	dbname   = "hacktiv8_employees"
)

var (
	db  *sql.DB
	err error
)

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected to database!")

	// CreateEmployee()
	// GetEmployee()
	// UpdateEmployee()
	DeleteEmployee()
}

type Employee struct {
	ID        int
	Full_name string
	Email     string
	Age       int
	Division  string
}

func CreateEmployee() {
	var employee = Employee{}

	sqlStatement := `
		INSERT INTO employees (full_name, email, age, division) 
		VALUES ($1, $2, $3, $4) 
		RETURNING *
	`

	err = db.QueryRow(sqlStatement, "Nur Sasongko", "kokonur1020@gmail.com", 21, "IT").
		Scan(&employee.ID, &employee.Full_name, &employee.Email, &employee.Age, &employee.Division)

	if err != nil {
		panic(err)
	}

	fmt.Printf("New employee data: %v has been created", employee)
}

func GetEmployee() {
	var results = []Employee{}

	sqlStatement := `SELECT * FROM employees`

	rows, err := db.Query(sqlStatement)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var employee = Employee{}

		err = rows.Scan(&employee.ID, &employee.Full_name, &employee.Email, &employee.Age, &employee.Division)

		if err != nil {
			panic(err)
		}

		results = append(results, employee)
	}

	fmt.Printf("All employees data: %v", results)
}

func UpdateEmployee() {
	sqlStatement := `
		UPDATE employees
		SET full_name = $2, email = $3, age = $4, division = $5
		WHERE id = $1
	`
	res, err := db.Exec(sqlStatement, 1, "Winawati", "myimouto@gmail.com", 17, "student")

	if err != nil {
		panic(err)
	}

	count, err := res.RowsAffected()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Updated %d employee data", count)
}

func DeleteEmployee() {
	sqlStatement := `
		DELETE FROM employees
		WHERE id = $1
	`
	res, err := db.Exec(sqlStatement, 1)

	if err != nil {
		panic(err)
	}

	count, err := res.RowsAffected()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Deleted %d employee data", count)
}
