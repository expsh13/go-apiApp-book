package services

import (
	"github.com/expsh13/go-apiApp-book/models"
	"github.com/expsh13/go-apiApp-book/repositories"
)

// ArticleDetailHandler で使うことを想定したサービス
// 指定 ID の記事情報を返却
func GetArticleService(articleID int) (models.Article, error) {
	// 1. repositories 層の関数 SelectArticleDetail で記事の詳細を取得
	db, err := connectDB()
	if err != nil {
		return models.Article{}, err
	}
	defer db.Close()
	article, err := repositories.SelectArticleDetail(db, articleID)
	if err != nil {
		return models.Article{}, err
	}
	// 2. repositories 層の関数 SelectCommentList でコメント一覧を取得
	commentList, err := repositories.SelectCommentList(db, articleID)
	if err != nil {
		return models.Article{}, err
	}
	// 3. 2 で得たコメント一覧を、1 で得た Article 構造体に紐付ける
	article.CommentList = append(article.CommentList, commentList...)
	return article, nil
}

// PostArticleHandler で使うことを想定したサービス
// 引数の情報をもとに新しい記事を作り、結果を返却
func PostArticleService(article models.Article) (models.Article, error) {
	// TODO : 実装
	db, err := connectDB()
	if err != nil {
		return models.Article{}, err
	}
	defer db.Close()
	newaArticle, err := repositories.InsertArticle(db, article)
	if err != nil {
		return models.Article{}, err
	}

	return newaArticle, nil
}

// ArticleListHandler で使うことを想定したサービス
// 指定 page の記事一覧を返却
func GetArticleListService(page int) ([]models.Article, error) {
	// TODO : 実装
	db, err := connectDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()
	articleArray, err := repositories.SelectArticleList(db, page)
	if err != nil {
		return nil, err
	}

	return articleArray, nil
}

// PostNiceHandler で使うことを想定したサービス
// 指定 ID の記事のいいね数を+1 して、結果を返却
func PostNiceService(article models.Article) (models.Article, error) {
	// TODO : 実装
	db, err := connectDB()
	if err != nil {
		return models.Article{}, err
	}
	defer db.Close()
	err = repositories.UpdateNiceNum(db, article.ID)
	if err != nil {
		return models.Article{}, err
	}

	return models.Article{
		ID:        article.ID,
		Title:     article.Title,
		Contents:  article.Contents,
		UserName:  article.UserName,
		NiceNum:   article.NiceNum + 1,
		CreatedAt: article.CreatedAt,
	}, nil
}
