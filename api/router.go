package api

import (
	"database/sql"
	"net/http"

	"github.com/expsh13/go-apiApp-book/controllers"
	"github.com/expsh13/go-apiApp-book/services"
	"github.com/gorilla/mux"
)

func NewRouter(db *sql.DB) *mux.Router {
	ser := services.NewMyAppService(db)
	aCon := controllers.NewArticleController(ser)
	cCon := controllers.NewCommentController(ser)
	// ルーティングとは「サーバーが受け取った HTTP リクエストを、どのハンドラに渡して処理させるのかを決める」操作のこと
	// ルーティング操作を担うものをルータ
	// gorilla/mux が用意しているルータを作成
	r := mux.NewRouter()
	r.HandleFunc("/hello", aCon.HelloHandler).Methods(http.MethodGet)

	r.HandleFunc("/article", aCon.PostArticleHandler).Methods(http.MethodPost)
	r.HandleFunc("/article/list", aCon.ArticleListHandler).Methods(http.MethodGet)
	r.HandleFunc("/article/{id:[0-9]+}", aCon.ArticleDetailHandler).Methods(http.MethodGet)
	r.HandleFunc("/article/nice", aCon.PostNiceHandler).Methods(http.MethodPost)

	r.HandleFunc("/comment", cCon.PostCommentHandler).Methods(http.MethodPost)
	// の対応付けが完了した状態のルータを返す
	return r
}
