package logic

import (
	"context"

	"mall/service/inventory/rpc/new/internal/svc"
	"mall/service/inventory/rpc/pb/inventory"

	"github.com/zeromicro/go-zero/core/logx"
)

type WmsProductDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewWmsProductDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *WmsProductDeleteLogic {
	return &WmsProductDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *WmsProductDeleteLogic) WmsProductDelete(in *inventory.WmsProductRemoveRequest) (*inventory.WmsProductRemoveResponse, error) {
	// todo: add your logic here and delete this line

	return &inventory.WmsProductRemoveResponse{}, nil
}
