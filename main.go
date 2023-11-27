package main

import (
	"log"
	"net/http"

	"github.com/expsh13/go-apiApp-book/handlers"
)

func main() {

	// ルーティングとは「サーバーが受け取った HTTP リクエストを、どのハンドラに渡して処理させるのかを決める」操作のこと
	// ルーティング操作を担うものをルータ
	http.HandleFunc("/hello", handlers.HelloHandler)
	http.HandleFunc("/article", handlers.PostArticleHandler)
	http.HandleFunc("/article/list", handlers.ArticleListHandler)
	http.HandleFunc("/article/1", handlers.ArticleDetailHandler)
	http.HandleFunc("/article/nice", handlers.PostNiceHandler)
	http.HandleFunc("/comment", handlers.PostCommentHandler)

	log.Println("server start at port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
