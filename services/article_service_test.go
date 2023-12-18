package services_test

import (
	"database/sql"
	"fmt"
	"os"
	"testing"

	"github.com/expsh13/go-apiApp-book/services"
	_ "github.com/go-sql-driver/mysql"
)

var aSer *services.MyAppService

func TestMain(m *testing.M) {
	// sql.DB 型を作る
	dbUser := "docker"
	dbPassword := "docker"
	dbDatabase := "sampledb"
	dbConn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?parseTime=true", dbUser, dbPassword, dbDatabase)
	db, err := sql.Open("mysql", dbConn)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// sql.DB 型からサービス構造体を作成
	aSer = services.NewMyAppService(db)
	// 個別のベンチマークテストの実行
	m.Run()
}

func BenchmarkGetArticleService(b *testing.B) {
	articleID := 1
	b.ResetTimer()
	// テスト関数の中でのベンチマークの取り方は「処理時間を測りたい関数・メソッドを複数回実行し、その平均を求める」という方法
	for i := 0; i < b.N; i++ {
		_, err := aSer.GetArticleService(articleID)
		if err != nil {
			b.Error(err)
			break
		}
	}
}
