package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"test/internal/logic"
	"test/internal/svc"
	"test/internal/types"
)

func TestHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Request
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewTestLogic(r.Context(), svcCtx)
		resp, err := l.Test(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
