package logic

import (
	"context"

	"mall/service/inventory/rpc/new/internal/svc"
	"mall/service/inventory/rpc/pb/inventory"

	"github.com/zeromicro/go-zero/core/logx"
)

type WarehouseCreateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewWarehouseCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *WarehouseCreateLogic {
	return &WarehouseCreateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *WarehouseCreateLogic) WarehouseCreate(in *inventory.WarehouseCreateRequest) (*inventory.WarehouseCreateResponse, error) {
	// todo: add your logic here and delete this line

	return &inventory.WarehouseCreateResponse{}, nil
}
