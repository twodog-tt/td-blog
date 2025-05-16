package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ ArticleCategoryModel = (*customArticleCategoryModel)(nil)

type (
	// ArticleCategoryModel is an interface to be customized, add more methods here,
	// and implement the added methods in customArticleCategoryModel.
	ArticleCategoryModel interface {
		articleCategoryModel
		withSession(session sqlx.Session) ArticleCategoryModel
	}

	customArticleCategoryModel struct {
		*defaultArticleCategoryModel
	}
)

// NewArticleCategoryModel returns a model for the database table.
func NewArticleCategoryModel(conn sqlx.SqlConn) ArticleCategoryModel {
	return &customArticleCategoryModel{
		defaultArticleCategoryModel: newArticleCategoryModel(conn),
	}
}

func (m *customArticleCategoryModel) withSession(session sqlx.Session) ArticleCategoryModel {
	return NewArticleCategoryModel(sqlx.NewSqlConnFromSession(session))
}
