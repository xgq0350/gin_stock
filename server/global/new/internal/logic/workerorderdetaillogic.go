package logic

import (
	"context"

	"mall/service/inventory/rpc/new/internal/svc"
	"mall/service/inventory/rpc/pb/inventory"

	"github.com/zeromicro/go-zero/core/logx"
)

type WorkerorderDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewWorkerorderDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *WorkerorderDetailLogic {
	return &WorkerorderDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *WorkerorderDetailLogic) WorkerorderDetail(in *inventory.NormalDetailRequest) (*inventory.WorkerorderDetailResponse, error) {
	// todo: add your logic here and delete this line

	return &inventory.WorkerorderDetailResponse{}, nil
}
