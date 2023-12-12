package services

import "github.com/expsh13/go-apiApp-book/models"

// 「そのインターフェースが指定しているメソッドを全て持つならば、どんな型でも受け付けられる
type MyAppServicer interface {
	PostArticleService(article models.Article) (models.Article, error)
	GetArticleListService(page int) ([]models.Article, error)
	GetArticleService(articleID int) (models.Article, error)
	PostNiceService(article models.Article) (models.Article, error)
	PostCommentService(comment models.Comment) (models.Comment, error)
}
