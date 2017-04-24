package main

import _ "github.com/denisenkom/go-mssqldb"
import "database/sql"
import "fmt"
import "log"
// import "os"
// import "bytes"
// import "bufio"

/*
	For more Go + SQL Server samples, please visit: aka.ms/go-sql
*/

var server = "localhost"
var port = 1433
var user = "sa"
var password = "Luis9000"

func main() {

	// tsql := FileToLines("./2 - insert_model.sql")

	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=FraudDetectionDemo", server, user, password, port)

	conn, err := sql.Open("mssql", connString)
	if err != nil {
		log.Fatal("Open connection failed:", err.Error())
	}
	fmt.Printf("Connected!\n")

	defer conn.Close()

	// Store model for native scoring
	// _, err = conn.Exec(tsql)
	// if err != nil {
	// 	fmt.Println("Error inserting the model: " + err.Error())
	// } else {
	// 	fmt.Printf("Saved fraud detection model!\n")
	// }
	fmt.Println("Scoring transactions:")
	ScoreTransactions(conn, "Fraud Detection Model (Native)", 0.8, "dbo.transaction_row_type")
	fmt.Println("")
}

// Score transactions
func ScoreTransactions(db *sql.DB, modelName string, fraudProbability float32, rowType string) (int, error) {
	fmt.Println("Declaring and inserting transaction rows")
	tsql := fmt.Sprintf("DECLARE @tx dbo.transaction_row_type; INSERT @tx VALUES (0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,GETDATE()),(0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,GETDATE());")
	_, err := db.Exec(tsql)
	if err != nil {
		fmt.Println("Error preparing the procedure: " + err.Error())
	}
	fmt.Println("Finished inserting transaction rows")

	stmt, err := db.Prepare("EXEC score_transactions ?, ?, ?;")
	if err != nil {
		fmt.Println("Error preparing the procedure: " + err.Error())
	} 

	rows, err := stmt.Query(modelName, fraudProbability, "@tx")
	if err != nil {
		fmt.Println("Error querying the procedure: " + err.Error())
	}

	defer stmt.Close()
	defer rows.Close()

	var count int = 0
	for rows.Next() {
		var transactionKey, transactionAmountUSD, transactionCurrencyCode, transactionDateTime, score string
		err := rows.Scan(&transactionKey, &transactionAmountUSD, &transactionCurrencyCode, &transactionDateTime, &score)
		if err != nil {
			fmt.Println("Error reading rows: " + err.Error())
			return -1, err
		}
		fmt.Printf("TransactionKey: %s, Amount: %s, Currency: %s, Timestamp: %s, Score: %s\n", transactionKey, transactionAmountUSD, transactionCurrencyCode, transactionDateTime, score)
		count++
	}
	return count, nil
}

// func FileToLines(filePath string) string {
// 	f, err := os.Open(filePath)
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer f.Close()

// 	var buffer bytes.Buffer
// 	scanner := bufio.NewScanner(f)
// 	for scanner.Scan() {
// 		buffer.WriteString(scanner.Text() + "\n")
// 	}
// 	if err := scanner.Err(); err != nil {
// 		fmt.Fprintln(os.Stderr, err)
// 	}

// 	return buffer.String()
// }
