package logic

import (
	"context"

	"mall/service/inventory/rpc/new/internal/svc"
	"mall/service/inventory/rpc/pb/inventory"

	"github.com/zeromicro/go-zero/core/logx"
)

type RepairListExLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRepairListExLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RepairListExLogic {
	return &RepairListExLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RepairListExLogic) RepairListEx(in *inventory.RepairListRequest) (*inventory.RepairListResponse, error) {
	// todo: add your logic here and delete this line

	return &inventory.RepairListResponse{}, nil
}
