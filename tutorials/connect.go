package main

import ( 
    _ "github.com/denisenkom/go-mssqldb"
    "database/sql"
    "context"
    "log"
    "fmt" 
)

// Replace with your own connection parameters
var server = "localhost"
var port = 1433
var user = "sa"
var password = "your_password"

var db *sql.DB

func main() {
    var err error

    // Create connection string
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d",
		server, user, password, port)
    
    // Create connection pool
	db, err = sql.Open("sqlserver", connString)
	if err != nil {
		log.Fatal("Error creating connection pool: " + err.Error())
	}
    log.Printf("Connected!\n")

    // Close the database connection pool after program executes
    defer db.Close()

    SelectVersion()
}

// Gets and prints SQL Server version 
func SelectVersion(){
    // Use background context 
    ctx := context.Background()

    // Ping database to see if it's still alive.
    // Important for handling network issues and long queries.
    err := db.PingContext(ctx)
	if err != nil {
		log.Fatal("Error pinging database: " + err.Error())
	}

	var result string

    // Run query and scan for result
	err = db.QueryRowContext(ctx, "SELECT @@version").Scan(&result)
	if err != nil {
        log.Fatal("Scan failed:", err.Error())
    }
    fmt.Printf("%s\n", result)
}