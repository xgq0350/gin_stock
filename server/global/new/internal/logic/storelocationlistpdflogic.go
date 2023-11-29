package logic

import (
	"context"

	"mall/service/inventory/rpc/new/internal/svc"
	"mall/service/inventory/rpc/pb/inventory"

	"github.com/zeromicro/go-zero/core/logx"
)

type StorelocationListPdfLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewStorelocationListPdfLogic(ctx context.Context, svcCtx *svc.ServiceContext) *StorelocationListPdfLogic {
	return &StorelocationListPdfLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *StorelocationListPdfLogic) StorelocationListPdf(in *inventory.StorelocationListRequest) (*inventory.NormalResponsePdf, error) {
	// todo: add your logic here and delete this line

	return &inventory.NormalResponsePdf{}, nil
}
