package handler

import (
	"net/http"

	"entrepreneurship/service/internal/logic"
	"entrepreneurship/service/internal/svc"
	"entrepreneurship/service/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func ServiceHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Request
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewServiceLogic(r.Context(), svcCtx)
		resp, err := l.Service(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
