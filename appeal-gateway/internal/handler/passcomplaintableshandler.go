package handler

import (
	"net/http"

	"appeal-gateway/internal/logic"
	"appeal-gateway/internal/svc"
	"appeal-gateway/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func PassComplainTablesHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ComplainPassReqest
		if err := httpx.ParseJsonBody(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}
		l := logic.NewPassComplainTablesLogic(r.Context(), svcCtx)
		resp, err := l.PassComplainTables(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
