package logic

import (
	"context"

	"mall/service/inventory/rpc/new/internal/svc"
	"mall/service/inventory/rpc/pb/inventory"

	"github.com/zeromicro/go-zero/core/logx"
)

type WmsProductDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewWmsProductDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *WmsProductDetailLogic {
	return &WmsProductDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *WmsProductDetailLogic) WmsProductDetail(in *inventory.WmsProductDetailRequest) (*inventory.WmsProductDetailResponse, error) {
	// todo: add your logic here and delete this line

	return &inventory.WmsProductDetailResponse{}, nil
}
