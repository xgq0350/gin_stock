package logic

import (
	"context"

	"mall/service/inventory/rpc/new/internal/svc"
	"mall/service/inventory/rpc/pb/inventory"

	"github.com/zeromicro/go-zero/core/logx"
)

type PalletPrintLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPalletPrintLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PalletPrintLogic {
	return &PalletPrintLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PalletPrintLogic) PalletPrint(in *inventory.PalletPrintRequest) (*inventory.PalletPrintResponse, error) {
	// todo: add your logic here and delete this line

	return &inventory.PalletPrintResponse{}, nil
}
