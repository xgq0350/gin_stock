package logic

import (
	"context"

	"mall/service/inventory/rpc/new/internal/svc"
	"mall/service/inventory/rpc/pb/inventory"

	"github.com/zeromicro/go-zero/core/logx"
)

type RepairDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRepairDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RepairDetailLogic {
	return &RepairDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RepairDetailLogic) RepairDetail(in *inventory.RepairDetailRequest) (*inventory.RepairDetailResponse, error) {
	// todo: add your logic here and delete this line

	return &inventory.RepairDetailResponse{}, nil
}
