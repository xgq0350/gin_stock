package logic

import (
	"context"

	"mall/service/inventory/rpc/new/internal/svc"
	"mall/service/inventory/rpc/pb/inventory"

	"github.com/zeromicro/go-zero/core/logx"
)

type StorelocationLocListPdfExLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewStorelocationLocListPdfExLogic(ctx context.Context, svcCtx *svc.ServiceContext) *StorelocationLocListPdfExLogic {
	return &StorelocationLocListPdfExLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *StorelocationLocListPdfExLogic) StorelocationLocListPdfEx(in *inventory.NormalIDListRequest) (*inventory.NormalResponsePdf, error) {
	// todo: add your logic here and delete this line

	return &inventory.NormalResponsePdf{}, nil
}
