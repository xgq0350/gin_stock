package logic

import (
	"context"

	"mall/service/inventory/rpc/new/internal/svc"
	"mall/service/inventory/rpc/pb/inventory"

	"github.com/zeromicro/go-zero/core/logx"
)

type WarehouseUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewWarehouseUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *WarehouseUpdateLogic {
	return &WarehouseUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *WarehouseUpdateLogic) WarehouseUpdate(in *inventory.WarehouseUpdateRequest) (*inventory.WarehouseUpdateResponse, error) {
	// todo: add your logic here and delete this line

	return &inventory.WarehouseUpdateResponse{}, nil
}
