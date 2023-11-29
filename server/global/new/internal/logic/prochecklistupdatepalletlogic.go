package logic

import (
	"context"

	"mall/service/inventory/rpc/new/internal/svc"
	"mall/service/inventory/rpc/pb/inventory"

	"github.com/zeromicro/go-zero/core/logx"
)

type ProchecklistUpdatePalletLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewProchecklistUpdatePalletLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProchecklistUpdatePalletLogic {
	return &ProchecklistUpdatePalletLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ProchecklistUpdatePalletLogic) ProchecklistUpdatePallet(in *inventory.PalletPrintRequest) (*inventory.ProchecklistUpdatePalletResponse, error) {
	// todo: add your logic here and delete this line

	return &inventory.ProchecklistUpdatePalletResponse{}, nil
}
