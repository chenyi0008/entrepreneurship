package logic

import (
	"context"

	"entrepreneurship/service/internal/svc"
	"entrepreneurship/service/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ServiceLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewServiceLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ServiceLogic {
	return &ServiceLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ServiceLogic) Service(req *types.Request) (resp *types.Response, err error) {
	resp = new(types.Response)
	resp.Message = req.Name

	return
}
