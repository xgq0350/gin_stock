package logic

import (
	"context"

	"mall/service/inventory/rpc/new/internal/svc"
	"mall/service/inventory/rpc/pb/inventory"

	"github.com/zeromicro/go-zero/core/logx"
)

type WaveOrderproductListPdfLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewWaveOrderproductListPdfLogic(ctx context.Context, svcCtx *svc.ServiceContext) *WaveOrderproductListPdfLogic {
	return &WaveOrderproductListPdfLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *WaveOrderproductListPdfLogic) WaveOrderproductListPdf(in *inventory.NormalIDListRequest) (*inventory.NormalResponsePdf, error) {
	// todo: add your logic here and delete this line

	return &inventory.NormalResponsePdf{}, nil
}
