package logic

import (
	"context"

	"mall/service/inventory/rpc/new/internal/svc"
	"mall/service/inventory/rpc/pb/inventory"

	"github.com/zeromicro/go-zero/core/logx"
)

type RepairCreateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRepairCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RepairCreateLogic {
	return &RepairCreateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RepairCreateLogic) RepairCreate(in *inventory.RepairCreateRequest) (*inventory.RepairCreateResponse, error) {
	// todo: add your logic here and delete this line

	return &inventory.RepairCreateResponse{}, nil
}
