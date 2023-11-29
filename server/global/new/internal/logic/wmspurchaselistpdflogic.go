package logic

import (
	"context"

	"mall/service/inventory/rpc/new/internal/svc"
	"mall/service/inventory/rpc/pb/inventory"

	"github.com/zeromicro/go-zero/core/logx"
)

type WmsPurchaseListPdfLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewWmsPurchaseListPdfLogic(ctx context.Context, svcCtx *svc.ServiceContext) *WmsPurchaseListPdfLogic {
	return &WmsPurchaseListPdfLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *WmsPurchaseListPdfLogic) WmsPurchaseListPdf(in *inventory.NormalIDListRequest) (*inventory.NormalResponsePdf, error) {
	// todo: add your logic here and delete this line

	return &inventory.NormalResponsePdf{}, nil
}
