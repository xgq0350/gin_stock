package logic

import (
	"context"

	"mall/service/inventory/rpc/new/internal/svc"
	"mall/service/inventory/rpc/pb/inventory"

	"github.com/zeromicro/go-zero/core/logx"
)

type PalletPrintDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPalletPrintDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PalletPrintDetailLogic {
	return &PalletPrintDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PalletPrintDetailLogic) PalletPrintDetail(in *inventory.PalletPrintRequest) (*inventory.PalletPrintListResponse, error) {
	// todo: add your logic here and delete this line

	return &inventory.PalletPrintListResponse{}, nil
}
