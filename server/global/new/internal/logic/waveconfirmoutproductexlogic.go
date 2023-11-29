package logic

import (
	"context"

	"mall/service/inventory/rpc/new/internal/svc"
	"mall/service/inventory/rpc/pb/inventory"

	"github.com/zeromicro/go-zero/core/logx"
)

type WaveConfirmOutProductExLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewWaveConfirmOutProductExLogic(ctx context.Context, svcCtx *svc.ServiceContext) *WaveConfirmOutProductExLogic {
	return &WaveConfirmOutProductExLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *WaveConfirmOutProductExLogic) WaveConfirmOutProductEx(in *inventory.ConfirmOutProductRequest) (*inventory.NormalResponse, error) {
	// todo: add your logic here and delete this line

	return &inventory.NormalResponse{}, nil
}
