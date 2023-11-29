package logic

import (
	"context"

	"mall/service/inventory/rpc/new/internal/svc"
	"mall/service/inventory/rpc/pb/inventory"

	"github.com/zeromicro/go-zero/core/logx"
)

type RepairListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRepairListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RepairListLogic {
	return &RepairListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RepairListLogic) RepairList(in *inventory.RepairListRequest) (*inventory.RepairListResponse, error) {
	// todo: add your logic here and delete this line

	return &inventory.RepairListResponse{}, nil
}
