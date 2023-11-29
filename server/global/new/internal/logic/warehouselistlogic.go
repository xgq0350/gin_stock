package logic

import (
	"context"

	"mall/service/inventory/rpc/new/internal/svc"
	"mall/service/inventory/rpc/pb/inventory"

	"github.com/zeromicro/go-zero/core/logx"
)

type WarehouseListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewWarehouseListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *WarehouseListLogic {
	return &WarehouseListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *WarehouseListLogic) WarehouseList(in *inventory.WarehouseListRequest) (*inventory.WarehouseListResponse, error) {
	// todo: add your logic here and delete this line

	return &inventory.WarehouseListResponse{}, nil
}
