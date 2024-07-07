package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "admin"
	dbname   = "postgres"
)

func main() {

	//http.HandleFunc("/", handler)
	//log.Fatal(http.ListenAndServe(":8080", nil))
	// Connection string
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	// Open connection to the database
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}
	defer db.Close()

	// Verify connection
	err = db.Ping()
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}
	fmt.Println("Successfully connected!")

	/* Create a table
		createTableSQL := `CREATE TABLE IF NOT EXISTS users (
	        id SERIAL PRIMARY KEY,
	        name TEXT,
	        age INT
	    );`
		_, err = db.Exec(createTableSQL)
		if err != nil {
			log.Fatalf("Error creating table: %v", err)
		}

		// Insert a new user
		insertSQL := `INSERT INTO users (name, age) VALUES ($1, $2)`
		_, err = db.Exec(insertSQL, "John Doe", 30)
		if err != nil {
			log.Fatalf("Error inserting new user: %v", err)
		}

		// Query the database
		var (
			id   int
			name string
			age  int
		)
		querySQL := `SELECT id, name, age FROM users WHERE name=$1`
		row := db.QueryRow(querySQL, "John Doe")
		err = row.Scan(&id, &name, &age)
		if err != nil {
			log.Fatalf("Error querying database: %v", err)
		}
		fmt.Printf("User found: %d, %s, %d\n", id, name, age)
	*/
}
