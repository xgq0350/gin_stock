package logic

import (
	"context"

	"mall/service/inventory/rpc/new/internal/svc"
	"mall/service/inventory/rpc/pb/inventory"

	"github.com/zeromicro/go-zero/core/logx"
)

type PalletPrintListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPalletPrintListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PalletPrintListLogic {
	return &PalletPrintListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PalletPrintListLogic) PalletPrintList(in *inventory.NormalListRequest) (*inventory.PalletPrintListResponse, error) {
	// todo: add your logic here and delete this line

	return &inventory.PalletPrintListResponse{}, nil
}
