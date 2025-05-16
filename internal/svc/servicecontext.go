package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/rest"
	"td-blog/internal/config"
	"td-blog/internal/middleware"
	"td-blog/model"
)

type ServiceContext struct {
	Config               config.Config
	Logging              rest.Middleware
	UserModel            model.UsersModel
	ArticlesModel        model.ArticlesModel
	CategoriesModel      model.CategoriesModel
	CommentsModel        model.CommentsModel
	ArticleCategoryModel model.ArticleCategoryModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:               c,
		Logging:              middleware.NewLoggingMiddleware().Handle,
		UserModel:            model.NewUsersModel(sqlx.NewMysql(c.DB.DataSource)),
		ArticlesModel:        model.NewArticlesModel(sqlx.NewMysql(c.DB.DataSource)),
		CategoriesModel:      model.NewCategoriesModel(sqlx.NewMysql(c.DB.DataSource)),
		CommentsModel:        model.NewCommentsModel(sqlx.NewMysql(c.DB.DataSource)),
		ArticleCategoryModel: model.NewArticleCategoryModel(sqlx.NewMysql(c.DB.DataSource)),
	}
}
