package logic

import (
	"context"

	"mall/service/inventory/rpc/new/internal/svc"
	"mall/service/inventory/rpc/pb/inventory"

	"github.com/zeromicro/go-zero/core/logx"
)

type RepairDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRepairDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RepairDeleteLogic {
	return &RepairDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RepairDeleteLogic) RepairDelete(in *inventory.RepairRemoveRequest) (*inventory.RepairRemoveResponse, error) {
	// todo: add your logic here and delete this line

	return &inventory.RepairRemoveResponse{}, nil
}
