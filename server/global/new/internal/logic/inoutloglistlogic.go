package logic

import (
	"context"

	"mall/service/inventory/rpc/new/internal/svc"
	"mall/service/inventory/rpc/pb/inventory"

	"github.com/zeromicro/go-zero/core/logx"
)

type InoutloglistLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewInoutloglistLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InoutloglistLogic {
	return &InoutloglistLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *InoutloglistLogic) Inoutloglist(in *inventory.NormalListRequest) (*inventory.InoutlogListResponse, error) {
	// todo: add your logic here and delete this line

	return &inventory.InoutlogListResponse{}, nil
}
