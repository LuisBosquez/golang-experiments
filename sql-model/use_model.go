package main

import _ "github.com/denisenkom/go-mssqldb"
import "database/sql"
import "log"
import "fmt"

/*
	For more Go + SQL Server samples, please visit: aka.ms/go-sql
 */

var server = "localhost"
var port = 1433
var user = "sa"
var password = "your_password"

func main() {
    connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d",
                                server, user, password, port)

    conn, err := sql.Open("mssql", connString)
    if err != nil {
        log.Fatal("Open connection failed:", err.Error())
    }
    fmt.Printf("Connected!\n")
    defer conn.Close()
    stmt, err := conn.Prepare("select @@version")
    row := stmt.QueryRow()
    var result string

    err = row.Scan(&result)
    if err != nil {
        log.Fatal("Scan failed:", err.Error())
    }
    fmt.Printf("%s\n", result)
}