package handler

import (
	"fmt"
	"net/http"

	"appeal-gateway/internal/logic"
	"appeal-gateway/internal/svc"
	"appeal-gateway/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetAppealListBySidHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AppealListRequest
		if err := httpx.ParseJsonBody(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}
		fmt.Println("111111111111")
		l := logic.NewGetAppealListBySidLogic(r.Context(), svcCtx)
		resp, err := l.GetAppealListBySid(&req)
		if err != nil {
			fmt.Println("11112111121111")
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}