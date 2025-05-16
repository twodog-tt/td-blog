package logic

import (
	"context"

	"td-blog/internal/svc"
	"td-blog/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListArticlesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListArticlesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListArticlesLogic {
	return &ListArticlesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListArticlesLogic) ListArticles(req *types.ListArticlesReq) (resp *types.ArticleListResp, err error) {
	// todo: add your logic here and delete this line

	return
}
