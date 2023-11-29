package logic

import (
	"context"

	"mall/service/inventory/rpc/new/internal/svc"
	"mall/service/inventory/rpc/pb/inventory"

	"github.com/zeromicro/go-zero/core/logx"
)

type WaveOrderlistLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewWaveOrderlistLogic(ctx context.Context, svcCtx *svc.ServiceContext) *WaveOrderlistLogic {
	return &WaveOrderlistLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *WaveOrderlistLogic) WaveOrderlist(in *inventory.NormalIDListRequest) (*inventory.WorkerorderListResponse, error) {
	// todo: add your logic here and delete this line

	return &inventory.WorkerorderListResponse{}, nil
}
