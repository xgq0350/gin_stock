package logic

import (
	"context"

	"mall/service/inventory/rpc/new/internal/svc"
	"mall/service/inventory/rpc/pb/inventory"

	"github.com/zeromicro/go-zero/core/logx"
)

type WmsProductListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewWmsProductListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *WmsProductListLogic {
	return &WmsProductListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *WmsProductListLogic) WmsProductList(in *inventory.WmsProductListRequest) (*inventory.WmsProductListResponse, error) {
	// todo: add your logic here and delete this line

	return &inventory.WmsProductListResponse{}, nil
}
