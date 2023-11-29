package logic

import (
	"context"

	"mall/service/inventory/rpc/new/internal/svc"
	"mall/service/inventory/rpc/pb/inventory"

	"github.com/zeromicro/go-zero/core/logx"
)

type WmsPurchaseproductListExLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewWmsPurchaseproductListExLogic(ctx context.Context, svcCtx *svc.ServiceContext) *WmsPurchaseproductListExLogic {
	return &WmsPurchaseproductListExLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *WmsPurchaseproductListExLogic) WmsPurchaseproductListEx(in *inventory.NormalIDListRequest) (*inventory.ProchecklistListResponse, error) {
	// todo: add your logic here and delete this line

	return &inventory.ProchecklistListResponse{}, nil
}
