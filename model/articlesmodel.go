package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ ArticlesModel = (*customArticlesModel)(nil)

type (
	// ArticlesModel is an interface to be customized, add more methods here,
	// and implement the added methods in customArticlesModel.
	ArticlesModel interface {
		articlesModel
		withSession(session sqlx.Session) ArticlesModel
	}

	customArticlesModel struct {
		*defaultArticlesModel
	}
)

// NewArticlesModel returns a model for the database table.
func NewArticlesModel(conn sqlx.SqlConn) ArticlesModel {
	return &customArticlesModel{
		defaultArticlesModel: newArticlesModel(conn),
	}
}

func (m *customArticlesModel) withSession(session sqlx.Session) ArticlesModel {
	return NewArticlesModel(sqlx.NewSqlConnFromSession(session))
}
