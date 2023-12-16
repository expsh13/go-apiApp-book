package controllers_test

import (
	"testing"

	"github.com/expsh13/go-apiApp-book/controllers"
	"github.com/expsh13/go-apiApp-book/controllers/testdata"

	_ "github.com/go-sql-driver/mysql"
)

// 1. テストに使うリソース (コントローラ構造体) を用意
var aCon *controllers.ArticleController

func TestMain(m *testing.M) {
	ser := testdata.NewServiceMock()
	aCon = controllers.NewArticleController(ser)
	m.Run()
}
