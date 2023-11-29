package logic

import (
	"context"

	"mall/service/inventory/rpc/new/internal/svc"
	"mall/service/inventory/rpc/pb/inventory"

	"github.com/zeromicro/go-zero/core/logx"
)

type WmsOrderproductListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewWmsOrderproductListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *WmsOrderproductListLogic {
	return &WmsOrderproductListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *WmsOrderproductListLogic) WmsOrderproductList(in *inventory.NormalIDListRequest) (*inventory.OrderproductListResponse, error) {
	// todo: add your logic here and delete this line

	return &inventory.OrderproductListResponse{}, nil
}
