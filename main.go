package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/expsh13/go-apiApp-book/controllers"
	"github.com/expsh13/go-apiApp-book/services"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
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

	ser := services.NewMyAppService(db)
	con := controllers.NewMyAppController(ser)

	// ルーティングとは「サーバーが受け取った HTTP リクエストを、どのハンドラに渡して処理させるのかを決める」操作のこと
	// ルーティング操作を担うものをルータ
	// gorilla/mux が用意しているルータを作成
	r := mux.NewRouter()

	r.HandleFunc("/hello", con.HelloHandler).Methods(http.MethodGet)
	r.HandleFunc("/article", con.PostArticleHandler).Methods(http.MethodPost)
	r.HandleFunc("/article/list", con.ArticleListHandler).Methods(http.MethodGet)
	r.HandleFunc("/article/{id:[0-9]+}", con.ArticleDetailHandler).Methods(http.MethodGet)
	r.HandleFunc("/article/nice", con.PostNiceHandler).Methods(http.MethodPost)

	r.HandleFunc("/comment", con.PostCommentHandler).Methods(http.MethodPost)

	log.Println("server start at port 8080")
	// http.ListenAndServe 関数にてサーバーを起動する際に、第二引数に gorilla/mux のルータを明示的に渡す
	log.Fatal(http.ListenAndServe(":8080", r))
}
