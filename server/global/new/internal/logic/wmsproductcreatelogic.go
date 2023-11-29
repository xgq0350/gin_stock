package logic

import (
	"context"

	"mall/service/inventory/rpc/new/internal/svc"
	"mall/service/inventory/rpc/pb/inventory"

	"github.com/zeromicro/go-zero/core/logx"
)

type WmsProductCreateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewWmsProductCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *WmsProductCreateLogic {
	return &WmsProductCreateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *WmsProductCreateLogic) WmsProductCreate(in *inventory.WmsProductCreateRequest) (*inventory.WmsProductCreateResponse, error) {
	// todo: add your logic here and delete this line

	return &inventory.WmsProductCreateResponse{}, nil
}
