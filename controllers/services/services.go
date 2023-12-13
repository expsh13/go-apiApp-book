package services

import "github.com/expsh13/go-apiApp-book/models"

// 「そのインターフェースが指定しているメソッドを全て持つならば、どんな型でも受け付けられる
// /article 関連を引き受けるサービス
type ArticleServicer interface {
	PostArticleService(article models.Article) (models.Article, error)
	GetArticleListService(page int) ([]models.Article, error)
	GetArticleService(articleID int) (models.Article, error)
	PostNiceService(article models.Article) (models.Article, error)
}

// /comment を引き受けるサービス
type CommentServicer interface {
	PostCommentService(comment models.Comment) (models.Comment, error)
}
