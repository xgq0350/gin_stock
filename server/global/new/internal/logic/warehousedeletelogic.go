package logic

import (
	"context"

	"mall/service/inventory/rpc/new/internal/svc"
	"mall/service/inventory/rpc/pb/inventory"

	"github.com/zeromicro/go-zero/core/logx"
)

type WarehouseDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewWarehouseDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *WarehouseDeleteLogic {
	return &WarehouseDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *WarehouseDeleteLogic) WarehouseDelete(in *inventory.WarehouseRemoveRequest) (*inventory.WarehouseRemoveResponse, error) {
	// todo: add your logic here and delete this line

	return &inventory.WarehouseRemoveResponse{}, nil
}
