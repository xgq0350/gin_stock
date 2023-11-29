package logic

import (
	"context"

	"mall/service/inventory/rpc/new/internal/svc"
	"mall/service/inventory/rpc/pb/inventory"

	"github.com/zeromicro/go-zero/core/logx"
)

type PalletPrintUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPalletPrintUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PalletPrintUpdateLogic {
	return &PalletPrintUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PalletPrintUpdateLogic) PalletPrintUpdate(in *inventory.PalletPrintRequest) (*inventory.PalletPrintListResponse, error) {
	// todo: add your logic here and delete this line

	return &inventory.PalletPrintListResponse{}, nil
}
