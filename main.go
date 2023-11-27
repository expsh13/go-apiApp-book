package main

import (
	"log"
	"net/http"

	"github.com/expsh13/go-apiApp-book/handlers"
	"github.com/gorilla/mux"
)

func main() {

	// ルーティングとは「サーバーが受け取った HTTP リクエストを、どのハンドラに渡して処理させるのかを決める」操作のこと
	// ルーティング操作を担うものをルータ
	// gorilla/mux が用意しているルータを作成
	r := mux.NewRouter()

	r.HandleFunc("/hello", handlers.HelloHandler).Methods(http.MethodGet)
	r.HandleFunc("/article", handlers.PostArticleHandler).Methods(http.MethodPost)
	r.HandleFunc("/article/list", handlers.ArticleListHandler).Methods(http.MethodGet)
	r.HandleFunc("/article/1", handlers.ArticleDetailHandler).Methods(http.MethodGet)
	r.HandleFunc("/article/nice", handlers.PostNiceHandler).Methods(http.MethodPost)
	r.HandleFunc("/comment", handlers.PostCommentHandler).Methods(http.MethodPost)

	log.Println("server start at port 8080")
	// http.ListenAndServe 関数にてサーバーを起動する際に、第二引数に gorilla/mux のルータを明示的に渡す
	log.Fatal(http.ListenAndServe(":8080", r))
}
