package logic

import (
	"context"

	"mall/service/inventory/rpc/new/internal/svc"
	"mall/service/inventory/rpc/pb/inventory"

	"github.com/zeromicro/go-zero/core/logx"
)

type WaveCreateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewWaveCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *WaveCreateLogic {
	return &WaveCreateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *WaveCreateLogic) WaveCreate(in *inventory.WaveCreateRequest) (*inventory.NormalResponse, error) {
	// todo: add your logic here and delete this line

	return &inventory.NormalResponse{}, nil
}
