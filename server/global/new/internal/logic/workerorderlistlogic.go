package logic

import (
	"context"

	"mall/service/inventory/rpc/new/internal/svc"
	"mall/service/inventory/rpc/pb/inventory"

	"github.com/zeromicro/go-zero/core/logx"
)

type WorkerorderListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewWorkerorderListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *WorkerorderListLogic {
	return &WorkerorderListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *WorkerorderListLogic) WorkerorderList(in *inventory.NormalStatusListRequest) (*inventory.WorkerorderListResponse, error) {
	// todo: add your logic here and delete this line

	return &inventory.WorkerorderListResponse{}, nil
}
