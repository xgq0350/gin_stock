package logic

import (
	"context"

	"mall/service/inventory/rpc/new/internal/svc"
	"mall/service/inventory/rpc/pb/inventory"

	"github.com/zeromicro/go-zero/core/logx"
)

type WarehouseDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewWarehouseDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *WarehouseDetailLogic {
	return &WarehouseDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *WarehouseDetailLogic) WarehouseDetail(in *inventory.WarehouseDetailRequest) (*inventory.WarehouseDetailResponse, error) {
	// todo: add your logic here and delete this line

	return &inventory.WarehouseDetailResponse{}, nil
}
