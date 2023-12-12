package routers

import (
	"net/http"

	"github.com/expsh13/go-apiApp-book/controllers"
	"github.com/gorilla/mux"
)

func NewRouter(con *controllers.MyAppController) *mux.Router {
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
	// の対応付けが完了した状態のルータを返す
	return r
}
