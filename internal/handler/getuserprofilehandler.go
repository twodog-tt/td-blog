package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"td-blog/internal/logic"
	"td-blog/internal/svc"
	"td-blog/internal/types"
)

func GetUserProfileHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.IDReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewGetUserProfileLogic(r.Context(), svcCtx)
		resp, err := l.GetUserProfile(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
