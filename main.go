package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/expsh13/go-apiApp-book/api"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func main() {
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
		log.Println("fail to connect DB")
		return
	}

	r := api.NewRouter(db)

	log.Println("server start at port 8080")
	// http.ListenAndServe 関数にてサーバーを起動する際に、第二引数に gorilla/mux のルータを明示的に渡す
	log.Fatal(http.ListenAndServe(":8080", r))
}
