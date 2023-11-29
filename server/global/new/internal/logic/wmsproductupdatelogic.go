package logic

import (
	"context"

	"mall/service/inventory/rpc/new/internal/svc"
	"mall/service/inventory/rpc/pb/inventory"

	"github.com/zeromicro/go-zero/core/logx"
)

type WmsProductUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewWmsProductUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *WmsProductUpdateLogic {
	return &WmsProductUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *WmsProductUpdateLogic) WmsProductUpdate(in *inventory.WmsProductUpdateRequest) (*inventory.WmsProductUpdateResponse, error) {
	// todo: add your logic here and delete this line

	return &inventory.WmsProductUpdateResponse{}, nil
}
