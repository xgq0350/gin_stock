package logic

import (
	"context"

	"mall/service/inventory/rpc/new/internal/svc"
	"mall/service/inventory/rpc/pb/inventory"

	"github.com/zeromicro/go-zero/core/logx"
)

type WavelistLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewWavelistLogic(ctx context.Context, svcCtx *svc.ServiceContext) *WavelistLogic {
	return &WavelistLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *WavelistLogic) Wavelist(in *inventory.NormalListRequest) (*inventory.WaveListResponse, error) {
	// todo: add your logic here and delete this line

	return &inventory.WaveListResponse{}, nil
}
