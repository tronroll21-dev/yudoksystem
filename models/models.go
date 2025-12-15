// models/models.go
package models

import (
	//"database/sql"
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql" // MySQL driver for database/sql
)

var db *sql.DB

func ConnectDB() error {

	// --- 2. Read parameters and interpolate DSN ---
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	/* データベース接続設定 */
	connStr := "%s:%s@tcp(%s:%s)/%s"
	var localErr error
	db, localErr = sql.Open("mysql", fmt.Sprintf(connStr, dbUser, dbPassword, dbHost, dbPort, dbName))
	if localErr != nil {
		return localErr
	}

	/* 接続確認（オプション） */
	if localErr = db.Ping(); localErr != nil {
		db.Close()
		return localErr
	}

	log.Println("Successfully connected to the database!")
	return nil
}

func InsertOrUpdateSalesRecord(input *DailyReportRaw) (record interface{}, found bool, mode string, err error) {
	// Dummy implementation for compilation; replace with real logic.
	// For example, check if record exists, update or insert accordingly.
	return input, false, "登録", nil
}
