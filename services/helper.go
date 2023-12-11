package services

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func connectDB() (*sql.DB, error) {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Printf("環境変数を読み込み出来ませんでした: %v", err)
	}
	dbUser := os.Getenv("USERNAME")
	dbPassword := os.Getenv("USERPASS")
	dbDatabase := os.Getenv("DATABASE")
	dbConn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?parseTime=true", dbUser, dbPassword, dbDatabase)
	db, err := sql.Open("mysql", dbConn)
	if err != nil {
		return nil, err
	}
	return db, nil
}
