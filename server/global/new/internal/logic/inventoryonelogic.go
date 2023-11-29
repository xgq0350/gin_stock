package logic

import (
	"context"

	"mall/service/inventory/rpc/new/internal/svc"
	"mall/service/inventory/rpc/pb/inventory"

	"github.com/zeromicro/go-zero/core/logx"
)

type InventoryOneLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewInventoryOneLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InventoryOneLogic {
	return &InventoryOneLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *InventoryOneLogic) InventoryOne(in *inventory.NormalDetailRequest) (*inventory.InventoryDetailResponse, error) {
	// todo: add your logic here and delete this line

	return &inventory.InventoryDetailResponse{}, nil
}
