package logic

import (
	"context"

	"mall/service/inventory/rpc/new/internal/svc"
	"mall/service/inventory/rpc/pb/inventory"

	"github.com/zeromicro/go-zero/core/logx"
)

type WaveProductlistLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewWaveProductlistLogic(ctx context.Context, svcCtx *svc.ServiceContext) *WaveProductlistLogic {
	return &WaveProductlistLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *WaveProductlistLogic) WaveProductlist(in *inventory.NormalIDListRequest) (*inventory.WaveProductListResponse, error) {
	// todo: add your logic here and delete this line

	return &inventory.WaveProductListResponse{}, nil
}
