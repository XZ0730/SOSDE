// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	"appeal-gateway/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/appealLeave",
				Handler: StudentAskforLeaveHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/appealLeave",
				Handler: GetAppealListBySidHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/stu/complain",
				Handler: ComplainToSupervisorHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/stu/complain",
				Handler: GetComplainTablesHandler(serverCtx),
			},
			{
				Method:  http.MethodDelete,
				Path:    "/stu/complain",
				Handler: PassComplainTablesHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/stu/appeal",
				Handler: PassAppealTableHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/hello",
				Handler: HelloHandler(serverCtx),
			},
		},
	)
}
