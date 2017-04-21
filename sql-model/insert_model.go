package main

import _ "github.com/denisenkom/go-mssqldb"
import "database/sql"
import "fmt"
import "os"
import "bytes"
import "log"
import "bufio"

/*
	For more Go + SQL Server samples, please visit: aka.ms/go-sql
 */


var server = "localhost"
var port = 1433
var user = "sa"
var password = "Luis9000"

func main() {

    tsql := FileToLines("./sql/model.sql")
    
    connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d",
                                server, user, password, port)

    conn, err := sql.Open("mssql", connString)
    if err != nil {
        log.Fatal("Open connection failed:", err.Error())
    }
    fmt.Printf("Connected!\n")

    defer conn.Close()

    stmt, err := conn.Prepare(tsql)
    row := stmt.QueryRow()
    var result string

    err = row.Scan(&result)
    if err != nil {
        log.Fatal("Scan failed:", err.Error())
    }
    
    fmt.Printf("%s\n", result)
}

func FileToLines(filePath string) string{
      f, err := os.Open(filePath)
      if err != nil {
        panic(err)
      }
      defer f.Close()

      var buffer bytes.Buffer
      scanner := bufio.NewScanner(f)
      for scanner.Scan() {
        buffer.WriteString(scanner.Text() + "\n")
      }
      if err := scanner.Err(); err != nil {
              fmt.Fprintln(os.Stderr, err)
      }

      return buffer.String()
}