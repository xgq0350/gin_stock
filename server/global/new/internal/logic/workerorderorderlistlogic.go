package logic

import (
	"context"

	"mall/service/inventory/rpc/new/internal/svc"
	"mall/service/inventory/rpc/pb/inventory"

	"github.com/zeromicro/go-zero/core/logx"
)

type WorkerorderOrderlistLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewWorkerorderOrderlistLogic(ctx context.Context, svcCtx *svc.ServiceContext) *WorkerorderOrderlistLogic {
	return &WorkerorderOrderlistLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *WorkerorderOrderlistLogic) WorkerorderOrderlist(in *inventory.NormalStatusListRequest) (*inventory.WorkerorderListResponse, error) {
	// todo: add your logic here and delete this line

	return &inventory.WorkerorderListResponse{}, nil
}
