package main

import (
	// "challenge-03/routers"
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

	// CreateBook()
	// GetBooks()
	// UpdateBook()
	DeleteBook()

	// var PORT = ":8080"

	// routers.StartServer().Run(PORT)
}

type Book struct {
	ID          int
	Title       string
	Author      string
	Description string
}

func CreateBook() {
	var book = Book{}

	sqlStatement := `
		INSERT INTO books (title, author, description) 
		VALUES ($1, $2, $3) 
		RETURNING *
	`

	err = db.QueryRow(sqlStatement, "The Little Book of Ikigai", "Ken Mogi", "Japanese Concept").
		Scan(&book.ID, &book.Title, &book.Author, &book.Description)

	if err != nil {
		panic(err)
	}

	fmt.Printf("New book data: %v has been created\n", book)
}

func GetBooks() {
	var results = []Book{}

	sqlStatement := `SELECT * FROM books`
	rows, err := db.Query(sqlStatement)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var book = Book{}

		err = rows.Scan(&book.ID, &book.Title, &book.Author, &book.Description)

		if err != nil {
			panic(err)
		}

		results = append(results, book)
	}

	fmt.Printf("All books data: %v\n", results)
}

func UpdateBook() {
	sqlStatement := `
		UPDATE books
		SET title = $2, author = $3, description = $4
		WHERE id = $1
	`
	res, err := db.Exec(sqlStatement, 1, "Buku Filsafat", "Winawati", "Buku Best Seller terbaru dari Winawati")

	if err != nil {
		panic(err)
	}

	count, err := res.RowsAffected()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Updated %d book data", count)
}

func DeleteBook() {
	sqlStatement := `
		DELETE FROM books
		WHERE id = $1
	`
	res, err := db.Exec(sqlStatement, 3)

	if err != nil {
		panic(err)
	}

	count, err := res.RowsAffected()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Deleted %d book data", count)
}
